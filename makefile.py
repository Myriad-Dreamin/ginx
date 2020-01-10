#!/usr/bin/env python3
import os, re, sys, shutil, json
from pymake import require_cls, oqs, entry, pipe


class Makefile:
    current_path = os.path.dirname(os.path.realpath(__file__))
    config_file_name = 'config.toml'
    context = dict()

    silent = False

    @classmethod
    def hello(cls, *_):
        print('minimum builder')

    @classmethod
    def pipe(cls, cmd, *_):
        if not cls.silent: print(cmd)
        pipe(cmd)

    @classmethod
    def generate(cls, path='./', match=None, *_):
        match = cls._gen_match(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                cls.generate(file, match)
            if os.path.isfile(file) and match.match(file):
                cls.pipe('go generate %s' % file)

    @staticmethod
    def _gen_match(match):
        if match is None:
            match = re.compile(r'^.*\.go$')
        if isinstance(match, str):
            match = re.compile(match)
        return match

    @classmethod
    @require_cls('up')
    def all(cls, *_):
        pass


if __name__ == '__main__':
    entry(Makefile, sys.argv[1:])
