#!/usr/bin/env python
# coding=utf-8

from __future__ import unicode_literals
from django.db import models
#import os
#os.environ['DJANGO_SETTINGS_MODULE'] = 'mymusic.settings'

class Singer(models.Model):
    """
        歌手信息表
    """

    singer_id = models.CharField('歌手ID', primary_key=True, max_length=100)
    singer_name = models.CharField('歌手名',max_length=100)


    def __str__(self):
        return self.singer_name
    def __unicode__(self):
        return self.singer_name



class Music_items(models.Model):
    """
        歌曲信息表
    """

    music_id = models.CharField('音乐ID', primary_key=True, max_length=100)
    music_url = models.CharField('歌曲 url',max_length=1024)
    music_name = models.CharField('歌曲名',max_length=1024)
    music_album = models.CharField('歌曲所属专辑',max_length=1024)
    music_mv = models.CharField('歌曲 mv',max_length=1024)
    music_artist = models.CharField('所属歌手',max_length=1024)
    #music_artist = models.ForeignKey('Singer', verbose_name='所属歌手',null=True,on_delete=models.SET_NULL)
    music_artist_id = models.CharField('歌手ID', max_length=100)

    def __str__(self):
        return self.music_name
    def __unicode__(self):
        return self.music_name


class Musiccomment(models.Model):
    """
        歌曲评论表
    """
    #music = models.ForeignKey('Music_items', verbose_name='歌曲 id',null=True,on_delete=models.SET_NULL)
    music_id = models.CharField('歌曲 id',max_length=100)
    comment_content = models.CharField('评论内容',max_length=10000)
    comment_user_id = models.CharField('评论用户 id',max_length=100)
    comment_username = models.CharField('评论用户名',max_length=100)
    comment_like_count = models.CharField('评论喜欢数',max_length=100)
    comment_id = models.CharField('评论 id',max_length=100)

    def __str__(self):
        return self.comment_content
    def __unicode__(self):
        return self.comment_content




