#!/usr/bin/env python
# -*- encoding=utf-8 -*-

import weakref
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
    """
    creates = []
    for t in ModelBase.metadata.sorted_tables:
        creates.append(str(CreateTable(t).compile(dialect=mysql.dialect()))+";")
    return '\n\n'.join(creates)

__all__ = ['model_registry', 'ModelBase', 'create_all']
