#!/usr/bin/env python
# coding=utf-8

import sys
reload(sys)
sys.setdefaultencoding('utf-8')
from bs4 import BeautifulSoup
import requests
import json
import re
from pickle import dump
from collections import OrderedDict
import re

top50_singer_ID_set=[]

def write_to_file(result):
    print type(result)
    f=open("./res","a")
    f.write(result)
    #dump(res,f)
    #for i in result:
        #f.write(str(i))
        #f.write('\n')
    f.write('\n')
    f.close()


def get_top_50():
    top50_singer_url='http://music.163.com/playlist?id=119712779'
    web_data=requests.get(top50_singer_url)
    #soup=BeautifulSoup(web_data.text,'lxml')
    soup=BeautifulSoup(web_data.text, "html.parser")
    #print soup.contents
    # 写入到文件
    #write_to_file(json.loads(soup.contents))

    R=soup.textarea.text #找到歌手ID所在的标签
    #R=soup.textarea.text #找到歌手ID所在的标签

    json_obj=json.loads(R)
    #print json.dumps(json_obj)
    #top50_singer_ID_set=[]
    for each in json_obj:

        singer_info = OrderedDict()
        a=json.dumps(each)
        res=a.decode("unicode-escape")
        #res= a.decode("unicode-escape").decode("unicode-escape")
        #print each
        #print res
        #write_to_file(res)
        singer_ID=each['artists'][0]['id']

        #print each['artists'][0]['name']

        # 将歌手信息存到 mysql
        singer_info["singer_id"] = each['artists'][0]['id']
        singer_info["singer_name"] = each['artists'][0]['name']

        top50_singer_ID_set.append(singer_ID)   #将排名前50的歌手的id存进一个列表

    # 列表去重
    top50_singer_ID = sorted(set(top50_singer_ID_set),key=top50_singer_ID_set.index)

    #print top50_singer_ID
    #write_to_file(top50_singer_ID_set)

def get_song_info(singer_ID):
    singer_url  = 'http://music.163.com/artist?id=' + str(singer_ID)
    #singer_url  = 'http://music.163.com/artist?id=8325'
    web_data=requests.get(singer_url)

    soup=BeautifulSoup(web_data.text,'lxml')
    #print soup.contents
    singer_name=soup.select("#artist-name")


    data=soup.select('div.n-artist')
    pattern1 = re.compile(r'<div.*?src="(.*?)".*?</div>', re.S)

    for item in data:
        final = re.findall(pattern1, str(item))
        if len(final) == 1:
            singer_pic = final[0]

    print singer_pic
    #singer_pic = singer_pic[0].get("src")


    #print type(singer_name)
    #for i in singer_name:
        #print i
    #res=''.join(singer_name)
    #print singer_name
    #each=json.loads(singer_name)
    #a=json.dumps(each)
    #res=singer_name.decode("unicode-escape")
    #print res


    singer_name=singer_name[0].get('title')

    r=soup.find('ul',{'class':'f-hide'}).find_all('a')

    r=(list(r))
    music_id_set=[]
    music_name_set=[]
    for each in r:
        song_name=each.text#print(each.text)
        music_name_set.append(song_name)

        song_id=each.attrs["href"]
        music_id_set.append(song_id[9:])

    #dic=dict(map(lambda x,y:[x,y],music_name_set,music_id_set))#将音乐名字和音乐id组成一个字典

    print "music_id_set: \n",music_id_set

if __name__ == '__main__':
    get_top_50()
    for singer_ID in top50_singer_ID_set:#依次将列表中的id代表的歌手的歌词下载下来
        get_song_info(singer_ID)
