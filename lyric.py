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

    json_obj=json.loads(R)
    print json.dumps(json_obj)
    #top50_singer_ID_set=[]
    for each in json_obj:
        a=json.dumps(each)
        #res= a.decode("unicode-escape").decode("unicode-escape")
        #print each
        res=a.decode("unicode-escape")
        #print res
        write_to_file(res)
        singer_ID=each['artists'][0]['id']
        top50_singer_ID_set.append(singer_ID)   #将排名前50的歌手的id存进一个列表

    # 列表去重
    top50_singer_ID = sorted(set(top50_singer_ID_set),key=top50_singer_ID_set.index)

    print top50_singer_ID
    #write_to_file(top50_singer_ID_set)

get_top_50()


def func(singer_ID1):#定义一个函数，通过一个歌手的id下载其最火的五十首歌的全部歌词

    singer_url  = 'http://music.163.com/artist?id=' + str(singer_ID1)
    web_data=requests.get(singer_url)

    soup=BeautifulSoup(web_data.text,'lxml')

    singer_name=soup.select("#artist-name")

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



    dic=dict(map(lambda x,y:[x,y],music_name_set,music_id_set))#将音乐名字和音乐id组成一个字典


    def get_lyric_by_music_id(music_id):#定义一个函数，通过音乐的id得到歌词
        lrc_url = 'http://music.163.com/api/song/lyric?' + 'id=' + str(music_id) + '&lv=1&kv=1&tv=-1'

        lyric=requests.get(lrc_url)
        json_obj=lyric.text
        #print(json_obj)
        j=json.loads(json_obj)
        #print(type(j))#打印出来j的类型是字典
        try:#部分歌曲没有歌词，这里引入一个异常
            lrc=j['lrc']['lyric']
            pat=re.compile(r'\[.*\]')
            lrc=re.sub(pat,"",lrc)
            lrc=lrc.strip()
            return lrc
        except KeyError as e:
            pass
    x=0
    for i in music_id_set:
        x=x+1


        print(x)
        top_50_lyric=get_lyric_by_music_id(i)

        f=open("music/%s.txt" % singer_name,"ab")#单个文件存储一个歌手的50首热门歌曲的歌词并以歌手的名字命名
        try:#引入异常
            f.write(top_50_lyric.encode('utf-8'))

            f.close()
        except AttributeError as e2:
            pass
if __name__ == '__main__':
    get_top_50()
    #for singer_ID in top50_singer_ID_set:#依次将列表中的id代表的歌手的歌词下载下来
        #singer_ID1=singer_ID
        #func(singer_ID1)
