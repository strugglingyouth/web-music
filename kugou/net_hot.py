#!/usr/bin/env python
# coding=utf-8
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
import requests
from bs4 import BeautifulSoup
import re
import json
import Mysql
from collections import OrderedDict

mysql = Mysql.Mysql()


# 保存到 mysql
def save_mysql(table, dict):
    mysql.insertData(table,dict)

def get_one_page(url):
    wb_data = requests.get(url)
    wb_data.encoding = wb_data.apparent_encoding
    return  wb_data.text
    #if wb_data.status_code == 200:
        #return wb_data.text
    #else:
        #return None


def parse_one_page(html):
    soup = BeautifulSoup(html, 'lxml')
    print "soup.content:",soup.contents
    data = soup.select('div.pc_temp_songlist ul li')
    #data = soup.select('div.ranklist-wrapper.clearfix div.mod-newsong ul.song-list li')
    #pattern1 = re.compile(r'<li.*?<div class="index">(.*?)</div>.*?title="(.*?)".*?title="(.*?)".*?</li>', re.S)
    #pattern2 = re.compile(r'<li.*?<div class="index">(.*?)</div>.*?title="(.*?)".*?target="_blank">(.*?)</a>', re.S)
    pattern1 = re.compile(r'<li.*?href="(.*?)".*?title="(.*?)".*?</li>', re.S)
    pattern2 = re.compile(r'<li.*?<div class="index">(.*?)</div>.*?title="(.*?)".*?target="_blank">(.*?)</a>', re.S)

    #data=json.loads(data)
    print data

    wants = []
    for item in data:
        #print item
        final = re.findall(pattern1, str(item))
        if len(final) == 1:
            wants.append(final[0])
            print final[0]
        else:
            other = re.findall(pattern2, str(item))
            wants.append(other[0])


    print "wants:",wants
    return wants


if __name__ == '__main__':
    #url = 'http://www.kugou.com/yy/rank/home/1-6666.html?from=rank'
    # url = 'http://www.kugou.com/yy/rank/home/1-6666.html'   #飙升榜
    url = 'http://www.kugou.com/yy/rank/home/1-23784.html'  # net hot
    html = get_one_page(url)

    #print "html:",html

    data = parse_one_page(html)


    music = OrderedDict()
    table_name = 'kugou_net_hot'

    for item in data:
        music["music_url"] = item[0]
        music["music_name"] = item[1]

        save_mysql(table_name,music)

