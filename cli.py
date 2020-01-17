import json
import os
import re
import subprocess
import urllib.request

import fire

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
    return subprocess.Popen(cmd_str, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT,
                            cwd=cwd).stdout.read()


def cmds(cmd_str, cwd=None):
    print(cmd_str)
    return cmd(cmd_str, cwd)


def pcmds(cmd_str, cwd=None, encoding='utf-8'):
    print(cmds(cmd_str, cwd).decode(encoding))

from collections import namedtuple


dependencies = namedtuple(
    'Dependencies',
    ['package_path',
     'magic_package_bin',
     'stringer_generator',
     'pure_object_template_router_path',
     'pure_object_template_service_path', # 5
     'model_db_object_file',
     'model_sp_object_file',
     'model_object_file',
     'object_template_service_path',
     'service_object_entry_file', # 10
     'service_provider_file',
     'server_init_service_file',
     'root_provider_file',
     'object_router_file',
     'db_core_file', # 15
     'model_provider_file',
     'server_init_db_file',
     'control_gen_file',
     'control_gen_main_file',
     ])(*
    ['github.com/Myriad-Dreamin/minimum-template',
     'github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path',
     'golang.org/x/tools/cmd/stringer',
     'template/control/pure-object/pure-object-router.go.template',
     'template/service/pure-object/', # 5
     'model/db-layer/object.go',
     'model/sp-layer/object.go',
     'model/object.go',
     'service/object/',
     'service/object.go', # 10
     'service/provider.go',
     'server/service.go',
     'control/router/provider.go',
     'control/router/object-router.go',
     'model/db-layer/core.go', # 15
     'model/sp-layer/provider.go',
     'server/db.go',
     'control/gen/object.go',
     'control/gen/generate.go',
     ]
)

#


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
        key_values = map(lambda x: x if x[1] else x, map(lambda x: x.split('=', 2), key_values))
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
        print('minimum-cli v0.3')

    def redeploy(self):
        pcmds("git pull")
        pcmds('%s cli.py make image' % MinimumCli.python_interpreter)
        pcmds('%s cli.py make down' % MinimumCli.python_interpreter)
        pcmds('%s cli.py make up' % MinimumCli.python_interpreter)

    def install(self):
        pcmds('go install %s' % dependencies.magic_package_bin)
        pcmds('go install %s' % dependencies.stringer_generator)
        self.fast_generate()

    def fmt(self):
        pcmds('go fmt %s/...' % dependencies.package_path)

    def create_service(self, object_name, placeholder, service_folder=None, router_template=None):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, router_template)
        self._create_service(object_name, placeholder, service_folder)

    def create_pure_service(self, object_name, placeholder):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, dependencies.pure_object_template_router_path)
        self._create_service(object_name, placeholder, dependencies.pure_object_template_service_path)

    def create_router(self, object_name, placeholder, __object_router_file=None):
        self._update_obj_vars(object_name, placeholder)
        self._create_router(object_name, placeholder, __object_router_file)

    def create_template(self, object_name, placeholder):
        self._update_obj_vars(object_name, placeholder)
        self._create_model(object_name, placeholder)
        self._create_service(object_name, placeholder)
        self._create_router(object_name, placeholder)

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

    def _create_service(self, object_name, placeholder, __object_service_folder=None):
        object_service_folder = 'service/' + self.m_snake_name + '/'
        object_service_entry_file = 'service/' + self.m_snake_name + '.go'
        object_control_gen_file = 'control/gen/' + self.m_snake_name + '.go'
        service_provider_file = dependencies.service_provider_file
        server_init_service_file = dependencies.server_init_service_file
        control_gen_main_file = dependencies.control_gen_main_file
        self.templates_to(__object_service_folder or dependencies.object_template_service_path, object_service_folder)
        self.template_to(dependencies.service_object_entry_file, object_service_entry_file)
        self.template_to(dependencies.control_gen_file, object_control_gen_file)

        self.replace(
            control_gen_main_file,
            '//instantiate',
            '//instantiate'
            '\n    %sCate := Describe%sService(v1)' % (self.camel, self.up_camel),
        )

        self.replace(
            control_gen_main_file,
            '//to files',
            '//to files'
            '\n    %sCate.ToFile("%s.go")' % (self.camel, self.m_snake_name),
        )

        self.replace(
            control_gen_main_file,
            'err := artisan.NewService(',
            'err := artisan.NewService('
            '\n        %sCate,' % self.camel,
        )

        self.replace(
            service_provider_file,
            'type Provider struct {',
            'type Provider struct {\n    %sService %sService' % (self.camel, self.up_camel),
        )

        self.replace(
            service_provider_file,
            'switch ss := service.(type) {',
            'switch ss := service.(type) {'
            '\n    case %sService:'
            '\n        s.%sService = ss'
            '\n        s.subControllers = append(s.subControllers, JustProvide(&ss))'
            '\n        return' % (self.up_camel, self.camel),
        )
        self.replace(
            server_init_service_file,
            'for _, serviceResult := range []serviceResult{',
            'for _, serviceResult := range []serviceResult{'
            '\n        {"%sService", functional.Decay(service.New%sService(srv.Module))},' % (self.camel, self.up_camel),
        )

    def _create_router(self, object_name, placeholder, __object_router_file=None):
        self._update_obj_vars(object_name, placeholder)
        object_router_file = 'control/router/' + self.m_snake_name + '-router.go'
        root_provider_file = dependencies.root_provider_file

        self.template_to(__object_router_file or dependencies.object_router_file, object_router_file)

        self.replace(
            root_provider_file,
            'objectRouter *ObjectRouter\n',
            'objectRouter *ObjectRouter\n'
            '\n    %sRouter *%sRouter' % (self.camel, self.up_camel),
        )

        self.replace(
            root_provider_file,
            'switch ss := router.(type) {',
            'switch ss := router.(type) {'
            '\n    case *%sRouter:'
            '\n        s.%sRouter = ss' % (self.up_camel, self.camel),
        )

    def _create_model(self, object_name, placeholder):
        self._update_obj_vars(object_name, placeholder)
        object_file = 'model/db-layer/' + self.m_snake_name + '.go'
        object_entry_file = 'model/' + self.m_snake_name + '.go'
        object_sp_file = 'model/sp-layer/' + self.m_snake_name + '.go'
        db_core_file = dependencies.db_core_file
        model_provider_file = dependencies.model_provider_file
        server_init_db_file = dependencies.server_init_db_file

        self.template_to(dependencies.model_db_object_file, object_file)
        self.template_to(dependencies.model_object_file, object_entry_file)
        self.template_to(dependencies.model_sp_object_file, object_sp_file)

        # todo

        self.replace(
            db_core_file,
            '//migrations',
            '//migrations'
            '\n        %s{}.migrate,' % self.up_camel,
            )

        self.replace(
            db_core_file,
            '//injections',
            '//injections'
            '\n        inject%sTraits,' % self.up_camel,
            )

        self.replace(
            model_provider_file,
            'objectDB *ObjectDB',
            'objectDB *ObjectDB'
            '\n    %sDB *%sDB' % (self.camel, self.up_camel),
            )

        self.replace(
            model_provider_file,
            'switch ss := database.(type) {',
            'switch ss := database.(type) {'
            '\n    case *%sDB:'
            '\n        s.%sDB = ss' % (self.up_camel, self.camel),
            )

        self.replace(
            server_init_db_file,
            'for _, dbResult := range []dbResult{',
            'for _, dbResult := range []dbResult{\n'
            '        {"%sDB", functional.Decay(model.New%sDB(srv.Module))},' % (self.camel, self.up_camel),
            )


minimum = MinimumCli()

if __name__ == '__main__':
    fire.Fire(MinimumCli)
