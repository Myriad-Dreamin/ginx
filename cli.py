
import fire
import shutil, json, os, subprocess, re, urllib.request
from makefile import Makefile
from pymake import entry

def to_camel_case(snake_str):
    components = snake_str.split('_')
    return components[0] + ''.join(x.title() for x in components[1:])


_qut_object_name = "'object'"
_object_name = 'object'
_camel = 'object'
_up_camel = 'Object'
_placeholder = 'oid'


def cmd(cmd_str, cwd=None):
    return subprocess.Popen(cmd_str, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, cwd=cwd).stdout.read()


def cmds(cmd_str, cwd=None):
    print(cmd_str)
    return cmd(cmd_str, cwd)


def pcmds(cmd_str, cwd=None, encoding='utf-8'):
    print(cmds(cmd_str, cwd).decode(encoding))


class MinimumCli:
    python_interpreter = 'python3'
    def __init__(self):
        self.qut_object_name = None
        self.object_name = None
        self.m_snake_name = None
        self.camel = None
        self.up_camel = None
        self.placeholder = None

    def make(self, *make_args):
        return entry(Makefile, make_args)

    def apply_context(self, *key_values, f=None, c=False):
        key_values = map(lambda x: x if x[1] else x,map(lambda x: x.split('=', 2), key_values))
        with open('.minimum-lib-env.json', 'r+') as f1:
            content = f1.read()
            context = json.loads('{}' if len(content) == 0 or c else content)
            if f is not None:
                with open(f, 'r+') as f2:
                    content = f2.read()
                    context.update(json.loads('{}' if len(content) == 0 else content))
            for key_value in key_values: context[key_value[0]] = key_value[1]
            f1.seek(0)
            f1.truncate()
            json.dump(context, f1, indent=4)

    def hello(self):
        print('minimum-cli v0.2')

    def get_cities(self):
        urllib.request.urlretrieve('https://raw.githubusercontent.com/wecatch/china_regions/master/json/city_object.json', 'city_object.json')

    def redeploy(self):
        pcmds("git pull")
        pcmds('%s cli.py make image' % MinimumCli.python_interpreter)
        pcmds('%s cli.py make down' % MinimumCli.python_interpreter)
        pcmds('%s cli.py make up' % MinimumCli.python_interpreter)

    def install(self):
        pcmds('go install github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path')
        pcmds('go install golang.org/x/tools/cmd/stringer')
        self.fast_generate()

    def fmt(self):
        pcmds('go fmt github.com/Myriad-Dreamin/minimum-lib/...')

    def create_service(self, object_name, placeholder, service_folder=None, router_template=None):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, router_template)
        self._create_service(object_name, placeholder, service_folder)

    def create_pure_service(self, object_name, placeholder):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, 'template/control/pure-object/pure-object-router.go.template')
        self._create_service(object_name, placeholder, 'template/pure-object/',
                             'template/control/pure-object/pure-object.go')

    def create_router(self, object_name, placeholder, __object_router_file=None):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, __object_router_file)

    def create_template(self, object_name, placeholder):
        self._update_obj_vars(object_name, placeholder)
        self.create_service(object_name, placeholder)
        self.create_router(object_name, placeholder)

        object_file = 'model/db-layer/' + self.m_snake_name + '.go'
        object_entry_file = 'model/' + self.m_snake_name + '.go'
        object_sp_file = 'model/sp-layer/' + self.m_snake_name + '.go'
        register_file = 'model/db-layer/register.go'
        model_provider_file = 'model/sp-layer/provider.go'

        server_init_db_file = 'server/db.go'

        self.template_to('model/db-layer/object.go', object_file)
        self.template_to('model/object.go', object_entry_file)
        self.template_to('model/sp-layer/object.go', object_sp_file)
        self.replace(
            register_file,
            'for _, migrate := range []func() error {',
            'for _, migrate := range []func() error {\n\t\t%s{}.migrate,' % self.up_camel,
        )

        self.replace(
            model_provider_file,
            'objectDB *ObjectDB',
            'objectDB *ObjectDB\n\t\t%sDB *%sDB,' % (self.camel, self.up_camel),
        )

        self.replace(
            model_provider_file,
            'switch ss := database.(type) {',
            'switch ss := database.(type) {\n\tcase *%sDB:\n\t\ts.%sDB = ss' % (self.up_camel, self.camel),
        )

        self.replace(
            server_init_db_file,
            'for _, dbResult := range []dbResult{',
            'for _, dbResult := range []dbResult{\n'
            '\t\t{"%sDB", functional.Decay(model.New%sDB(srv.Module))},' % (self.camel, self.up_camel),
        )

    def template_to(self, src, dst):
        # shutil.copyfile(src, dst)
        with open(src, 'r+') as obj_template:
            source = obj_template.read()
            source = source.replace(_qut_object_name, self.qut_object_name)
            source = source.replace(_camel, self.camel)
            source = source.replace(_up_camel, self.up_camel)
            source = source.replace(_placeholder, self.placeholder)
            with open(dst, 'w+') as obj_dump:
                obj_dump.truncate()
                obj_dump.write(source)

    def templates_to(self, src_path, dst_path):
        if not os.path.exists(dst_path):
            os.makedirs(dst_path)
        for file in os.listdir(src_path):
            dst_file = os.path.join(dst_path, file)
            file = os.path.join(src_path, file)
            if os.path.isdir(file):
                self.templates_to(file, dst_file)
            if os.path.isfile(file):
                self.template_to(file, dst_file)

    def test(self):
        self.fast_generate()
        pcmds('go test -v', cwd='./test')

    def generate(self, path='./', match=None):
        match = self._gen_match(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                self.generate(file, match)
            if os.path.isfile(file) and match.match(file):
                pcmds('go generate %s' % file)

    def fast_generate(self, path='./', match=None):
        match = self._gen_match(match)
        if isinstance(match, str):
            match = re.compile(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                self.fast_generate(file, match)
            if os.path.isfile(file) and match.match(file):
                with open(file, 'r', encoding='utf-8') as go_file:
                    for line in go_file.readlines():
                        if line.startswith('//go:generate '):
                            pcmds(line[len('//go:generate '):], cwd=path)

    def replace(self, file_name, old_str, new_str):
        with open(file_name, 'r+') as f:
            source = f.read()
            f.seek(0)
            f.truncate()
            f.write(source.replace(old_str, new_str))

    @staticmethod
    def _gen_match(match):
        if match is None:
            match = re.compile(r'^.*\.go$')
        if isinstance(match, str):
            match = re.compile(match)
        return match

    def _update_obj_vars(self, object_name, placeholder=None):
        self.object_name = object_name
        self.m_snake_name = self.object_name.replace('_', '-')
        self.qut_object_name = "'" + self.object_name + "'"
        self.camel = to_camel_case(object_name)
        self.up_camel = self.camel.title()
        self.placeholder = placeholder

    def _create_service(self, object_name, placeholder, __object_service_folder=None, __control_service_template=None):
        object_service_folder = 'service/' + self.m_snake_name + '/'
        object_service_entry_file = 'service/' + self.m_snake_name + '.go'
        object_interface_entry_file = 'control/' + self.m_snake_name + '.go'
        service_provider_file = 'service/provider.go'

        self.templates_to(__object_service_folder or 'service/object/', object_service_folder)
        self.template_to('service/object.go', object_service_entry_file)
        self.template_to(__control_service_template or 'control/interface.go', object_interface_entry_file)

        self.replace(
            service_provider_file,
            'type Provider struct {',
            'type Provider struct {\n\t%sService %sService' % (self.camel, self.up_camel),
        )

        self.replace(
            service_provider_file,
            'switch ss := service.(type) {',
            'switch ss := service.(type) {'
            '\n\tcase %sService:'
            '\n\t\ts.%sService = ss'
            '\n\t\ts.subControllers = append(s.subControllers, JustProvide(&ss))' % (self.up_camel, self.camel),
        )
        server_init_service_file = 'server/service.go'
        self.replace(
            server_init_service_file,
            'for _, serviceResult := range []serviceResult{',
            'for _, serviceResult := range []serviceResult{\n'
            '\t\t{"%sService", functional.Decay(service.New%sService(srv.Module))},' % (self.camel, self.up_camel),
        )

    def _create_router(self, object_name, placeholder, __object_router_file=None):
        self._update_obj_vars(object_name, placeholder)
        object_router_file = 'control/router/' + self.m_snake_name + '-router.go'
        root_provider_file = 'control/router/provider.go'

        self.template_to(__object_router_file or 'control/router/object-router.go', object_router_file)

        self.replace(
            root_provider_file,
            'objectRouter *ObjectRouter\n',
            'objectRouter *ObjectRouter\n\n\t%sRouter *%sRouter' % (self.camel, self.up_camel),
        )
        self.replace(
            root_provider_file,
            'switch ss := router.(type) {',
            'switch ss := router.(type) {\n\tcase *%sRouter:\n\t\ts.%sRouter = ss' % (self.up_camel, self.camel),
        )

minimum = MinimumCli()

if __name__ == '__main__':
    fire.Fire(MinimumCli)