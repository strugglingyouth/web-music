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
    singer_pic = models.CharField('歌手名',max_length=1024)

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


class User(models.Model):
    """
        用户表
    """
    user_name = models.CharField('用户名',primary_key=True,max_length=100)
    user_password = models.CharField('密码',max_length=100)
    user_nick_name = models.CharField('昵称',max_length=100)
    user_birth = models.CharField('生日',max_length=100)
    user_sex = models.CharField('性别',max_length=100)
    user_intro = models.CharField('个人简介',max_length=10024)
    user_open = models.PositiveIntegerField('是否公开个人信息(粉丝，所有人，保密)(0,1,-1)',default=1)
    user_list_open = models.BooleanField('是否公开收藏信息',default=True)

    def __str__(self):
        return self.user_nick_name
    def __unicode__(self):
        return self.user_nick_name


class Upload(models.Model):
    """
        用户上传表
    """
    upload_user_name = models.CharField('用户名',max_length=100)
    upload_music_name = models.CharField('上传歌曲名',max_length=100)
    upload_open = models.BooleanField('是否公开',default=False)
    upload_date = models.DateTimeField('创建时间', auto_now_add=True)

    class Meta:
        ordering = ['-upload_date']

    def __str__(self):
        return self.upload_user_name
    def __unicode__(self):
        return self.upload_user_name


class Follow(models.Model):
    """
        关注表
    """


    follow_id = models.CharField('被关注id',primary_key=True,max_length=100)
    follow_user_id = models.CharField('关注者id',max_length=100)

    def __str__(self):
        return self.follow_id
    def __unicode__(self):
        return self.follow_id



class My_list(models.Model):
    """
        歌单表
    """

    my_list_user_id = models.CharField('创建用户id',max_length=100)
    my_list_name = models.CharField('歌单名',max_length=100)
    my_list_count = models.CharField('歌单歌曲数',max_length=100)
    my_list_open = models.BooleanField('是否公开',default=True)
    create_date = models.DateTimeField('创建时间', auto_now_add=True)

    def __str__(self):
        return self.my_list_name
    def __unicode__(self):
        return self.my_list_name



class My_list_to_music(models.Model):
    """
        歌单关联到歌曲表
    """

    my_list_id = models.CharField('歌单id',primary_key=True,max_length=100)
    music_id = models.CharField('歌曲id',max_length=100)

    def __str__(self):
        return self.my_list_id
    def __unicode__(self):
        return self.my_list_id

class Collection(models.Model):
    """
        用户收藏歌单表
    """

    user_id = models.CharField('用户id',primary_key=True,max_length=100)
    my_list_id = models.CharField('歌单id',max_length=100)


    def __str__(self):
        return self.user_id
    def __unicode__(self):
        return self.user_id











