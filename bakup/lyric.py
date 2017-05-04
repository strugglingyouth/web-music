#!/usr/bin/env python
# coding=utf-8
 1 # -*- coding:utf-8 -*-
 2 import requests
 3 from bs4 import BeautifulSoup
 4 import json
 5 import re
 6 top50_singer_url=‘http://music.163.com/playlist?id=119712779‘
 7 web_data=requests.get(top50_singer_url)
 8 soup=BeautifulSoup(web_data.text,‘lxml‘)
 9
10 R=soup.textarea.text#找到歌手ID所在的标签
11
12 json_obj=json.loads(R)
13 top50_singer_ID_set=[]
14 for each in json_obj:
15     singer_ID=each[‘artists‘][0][‘id‘]
16     top50_singer_ID_set.append(singer_ID)#将排名前50的歌手的id存进一个列表
17
18
19 def func(singer_ID1):#定义一个函数，通过一个歌手的id下载其最火的五十首歌的全部歌词
20
21
22     from bs4 import BeautifulSoup
23     singer_url  = ‘http://music.163.com/artist?id=‘ + str(singer_ID1)
24     web_data=requests.get(singer_url)
25     soup=BeautifulSoup(web_data.text,‘lxml‘)
26     singer_name=soup.select("#artist-name")
27
28     singer_name=singer_name[0].get(‘title‘)
29
30     r=soup.find(‘ul‘,{‘class‘:‘f-hide‘}).find_all(‘a‘)
31     r=(list(r))
32     music_id_set=[]
33     music_name_set=[]
34     for each in r:
35         song_name=each.text#print(each.text)
36         music_name_set.append(song_name)
37
38         song_id=each.attrs["href"]
39         music_id_set.append(song_id[9:])
40
41
42
43     dic=dict(map(lambda x,y:[x,y],music_name_set,music_id_set))#将音乐名字和音乐id组成一个字典
44
45
46     from bs4 import BeautifulSoup
47     def get_lyric_by_music_id(music_id):#定义一个函数，通过音乐的id得到歌词
48         lrc_url = ‘http://music.163.com/api/song/lyric?‘ + ‘id=‘ + str(music_id) + ‘&lv=1&kv=1&tv=-1‘
49
50         lyric=requests.get(lrc_url)
51         json_obj=lyric.text
52         #print(json_obj)
53         j=json.loads(json_obj)
54         #print(type(j))#打印出来j的类型是字典
55         try:#部分歌曲没有歌词，这里引入一个异常
56             lrc=j[‘lrc‘][‘lyric‘]
57             pat=re.compile(r‘\[.*\]‘)
58             lrc=re.sub(pat,"",lrc)
59             lrc=lrc.strip()
60             return lrc
61         except KeyError as e:
62             pass
63     x=0
64     for i in music_id_set:
65         x=x+1
66
67
68         print(x)
69         top_50_lyric=get_lyric_by_music_id(i)
70
71         f=open("F:/projects/scrapy/%s.txt" % singer_name,"ab")#单个文件存储一个歌手的50首热门歌曲的歌词并以歌手的名字命名
72         try:#引入异常
73             f.write(top_50_lyric.encode(‘utf-8‘))
74
75             f.close()
76         except AttributeError as e2:
77             pass
78 for singer_ID in top50_singer_ID_set:#依次将列表中的id代表的歌手的歌词下载下来
79     singer_ID1=singer_ID
80     func(singer_ID1)
