from __future__ import unicode_literals

from django.db import models


class Music_items(models.Model):
    """

    """

    name = models.CharField('类名',max_length=20)
    created_time = models.DateTimeField('创建时间',auto_now_add=True)
    last_modified_time = models.DateTimeField('修改时间',auto_now=True)

    def __str__(self):
        return self.name
    def __unicode__(self):
        return self.name




