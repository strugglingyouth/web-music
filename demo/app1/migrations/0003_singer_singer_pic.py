# -*- coding: utf-8 -*-
# Generated by Django 1.9.6 on 2017-05-07 13:14
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('app1', '0002_music_items_music_artist_id'),
    ]

    operations = [
        migrations.AddField(
            model_name='singer',
            name='singer_pic',
            field=models.CharField(default=0, max_length=1024, verbose_name='\u6b4c\u624b\u540d'),
            preserve_default=False,
        ),
    ]
