
import fire
import shutil
import os
import subprocess
import re


def to_camel_case(snake_str) -> str:
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
    def __init__(self):
        self.x = 1
        self.qut_object_name = None
        self.object_name = None
        self.m_snake_name = None
        self.camel = None
        self.up_camel = None
        self.placeholder = None

    def hello(self):
        print('hello', self.x)

    def install(self):
        pcmds('go install github.com/Myriad-Dreamin/minimum-lib/generate/minimum-attach-file')

    def create_template(self, object_name: str, placeholder):
        self.object_name = object_name
        self.m_snake_name = self.object_name.replace('_', '-')
        self.qut_object_name = "'" + self.object_name + "'"
        self.camel = to_camel_case(object_name)
        self.up_camel = self.camel.title()
        self.placeholder = placeholder

        object_file = 'model/db-layer/' + self.m_snake_name + '.go'
        object_entry_file = 'model/' + self.m_snake_name + '.go'
        object_service_entry_file = 'service/' + self.m_snake_name + '.go'
        object_sp_file = 'model/sp-layer/' + self.m_snake_name + '.go'
        object_service_folder = 'service/' + self.m_snake_name + '/'
        object_router_file = 'router/' + self.m_snake_name + '-router.go'
        register_file = 'model/db-layer/register.go'
        model_provider_file = 'model/sp-layer/provider.go'
        service_provider_file = 'service/provider.go'
        root_provider_file = 'router/provider.go'
        server_init_db_file = 'server/db.go'
        server_init_service_file = 'server/service.go'

        self.template_to('router/object-router.go', object_router_file)
        self.template_to('model/db-layer/object.go', object_file)
        self.template_to('model/object.go', object_entry_file)
        self.template_to('model/sp-layer/object.go', object_sp_file)
        self.templates_to('service/object/', object_service_folder)
        self.template_to('service/object.go', object_service_entry_file)
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
            service_provider_file,
            'type Provider struct {',
            'type Provider struct {\n\t%sService *%sService' % (self.camel, self.up_camel),
        )

        self.replace(
            service_provider_file,
            'switch ss := service.(type) {',
            'switch ss := service.(type) {'
            '\n\tcase *%sService:'
            '\n\t\ts.%sService = ss'
            '\n\t\ts.subControllers = append(s.subControllers, JustProvide(&ss))' % (self.up_camel, self.camel),
        )

        self.replace(
            root_provider_file,
            'objectRouter *ObjectRouter',
            'objectRouter *ObjectRouter\n\t%sRouter *%sRouter' % (self.camel, self.up_camel),
        )

        self.replace(
            root_provider_file,
            'switch ss := router.(type) {',
            'switch ss := router.(type) {\n\tcase *%sRouter:\n\t\ts.%sRouter = ss' % (self.up_camel, self.camel),
        )

        self.replace(
            server_init_db_file,
            'for _, dbResult := range []dbResult{',
            'for _, dbResult := range []dbResult{\n'
            '\t\t{"%sDB", traits.Decay(model.New%sDB(srv.Logger, srv.cfg))},' % (self.camel, self.up_camel),
        )

        self.replace(
            server_init_service_file,
            'for _, serviceResult := range []serviceResult{',
            'for _, serviceResult := range []serviceResult{\n'
            '\t\t{"%sService", traits.Decay(service.New%sService('
            'srv.Logger, srv.DatabaseProvider, srv.cfg))},' % (self.camel, self.up_camel),
        )

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

    @staticmethod
    def _gen_match(match):
        if match is None:
            match = re.compile(r'^.*\.go$')
        if isinstance(match, str):
            match = re.compile(match)
        return match

    def generate(self, path='./', match=None):
        match = self._gen_match(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                self.generate(file, match)
            if os.path.isfile(file) and match.match(file):
                pcmds('go generate %s' % file)

    def test(self):
        self.fast_generate()
        pcmds('go test -v', cwd='./test')

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


if __name__ == '__main__':
    fire.Fire(MinimumCli)

