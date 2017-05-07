#!/usr/bin/env python
# coding=utf-8
from bs4 import BeautifulSoup
from urllib2 import urlopen, Request, URLError, HTTPError
import time


def make_soup(url):
    # """打开指定url 获取BeautifulSoup对象"""
    try:
        req = Request(url)
        response = urlopen(req)
        html = response.read()
    except URLError, e:
        if hasattr(e, 'code'):
            print '错误码: ', e.code, ',无法完成请求.'
        elif hasattr(e, 'reason'):
            print '请求失败: ', e.reason, '无法连接服务器'
    else:
        return BeautifulSoup(html,'lxml')


def get_music(b_soup, sel):
    #""" 获取歌曲榜单"""
    main_div = b_soup.select(sel)[0]
    # 获取类别列表
    sum_category = main_div.select('p > strong > a[title]')[0].string
    titles = [sum_category+' '+a.string for a in main_div.select('p > span > a[title]')]
    index = 0
    song_dict = {}
    # 逐个解析下层的歌单并加入类别:歌单 字典对象
    for div in main_div.find_all('div', recursive=False):   # 这里我们不能递归查找
        part = div.find_all('span', class_='text')
        if part:
            song_dict[titles[index]] = part
            index += 1
    return song_dict


BASE_URL = 'http://www.kugou.com/'

#这是酷狗首页榜单的div选择器
#如果页面变动 需要更改此处
DIV_LIST = [
    'div#single0',   # 推荐歌曲部分div
    'div.clear_fix.hot_top_10',  # 热榜top10部分div
    'div.clear_fix.hot_global.hot_top_10'  # 全球热榜部分div
]


def main():
    soup = make_soup(BASE_URL)
    if soup is None:
        print '抱歉，无法完成抽取任务，即将退出...'
        exit()
    print '获取时间: '+time.strftime("%Y-%m-%d %H:%M:%S")
    for k in DIV_LIST:   # 从歌单div逐个解析
        for category, items in get_music(soup, k).iteritems():  # 打印类别:歌单字典对象内容
            print '*'*20+category+'*'*30
            count = 1
            for song in items:
                print count, song.string
                count += 1
    print '*'*60
    print '获取歌单结束'

if __name__ == "__main__":
    main()

