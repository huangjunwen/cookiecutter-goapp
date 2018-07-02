#!/usr/bin/env python
# -*- encoding=utf-8 -*-

import os
from os.path import dirname

__all__ = []

# 自动将该目录下所有 module 导入，并将其内容带到 globals 中，每个 module 应当设置 __all__
# https://stackoverflow.com/questions/1057431/how-to-load-all-modules-in-a-folder/1057765#1057765
# https://stackoverflow.com/questions/21221358/python-how-to-import-all-methods-and-attributes-from-a-module-dynamically/21221452#21221452
for f in os.listdir(dirname(__file__)):
    if not f.endswith('.py') or f == '__init__.py':
        continue
    module = __import__(f[:-3], locals(), globals())

    module_dict = module.__dict__
    try:
        names = module.__all__
    except AttributeError:
        names = [name for name in module_dict if not name.startswith('_')]
    globals().update({name: module_dict[name] for name in names})

    __all__.extend(names)
