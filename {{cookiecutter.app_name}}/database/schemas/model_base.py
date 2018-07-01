#!/usr/bin/env python
# -*- encoding=utf-8 -*-

import weakref
from sqlalchemy import Column, func
from sqlalchemy import Integer, DateTime
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

__all__ = ['model_registry', 'ModelBase']
