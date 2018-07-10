#!/usr/bin/env python
# -*- encoding=utf-8 -*-

import weakref
from StringIO import StringIO

import sqlalchemy as sa
from sqlalchemy import Column, func
from sqlalchemy import Integer, DateTime
from sqlalchemy.schema import CreateTable
from sqlalchemy.dialects import mysql
from sqlalchemy.ext.declarative import declarative_base


class ModelBaseObject(object):
    __table_args__ = {
        'mysql_engine': 'InnoDB',
        'mysql_charset': 'utf8mb4',
        'mysql_collate': 'utf8mb4_bin',
    }

    # 数值 id
    id = Column(Integer, primary_key=True)

    # 创建时间
    created_at = Column(DateTime, server_default=func.now(), nullable=False)


# model 名 -> model
model_registry = weakref.WeakValueDictionary()

# model 基类
ModelBase = declarative_base(
    cls=ModelBaseObject, class_registry=model_registry)

def create_all():
    """
    返回全部表的 create table
    ref: https://gist.github.com/eirnym/d9079e5eee380450f464
    """
    out = StringIO()
    def dump(sql, *args, **kws):
        out.write('{};\n'.format(str(sql.compile(dialect=dump.dialect)).strip()))

    engine = sa.create_engine("mysql://", strategy='mock', executor=dump)
    dump.dialect = engine.dialect

    ModelBase.metadata.create_all(engine)

    return out.getvalue()

__all__ = ['model_registry', 'ModelBase', 'create_all']
