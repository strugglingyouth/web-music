-- MySQL dump 10.16  Distrib 10.1.22-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: kugou
-- ------------------------------------------------------
-- Server version	10.1.21-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `kugou_biaoshengbang`
--

DROP TABLE IF EXISTS `kugou_biaoshengbang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kugou_biaoshengbang` (
  `music_id` int(11) NOT NULL AUTO_INCREMENT,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`music_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kugou_biaoshengbang`
--

LOCK TABLES `kugou_biaoshengbang` WRITE;
/*!40000 ALTER TABLE `kugou_biaoshengbang` DISABLE KEYS */;
INSERT INTO `kugou_biaoshengbang` VALUES (1,'http://www.kugou.com/song/fsxhb15.html','庄心妍 - 我不相信'),(2,'http://www.kugou.com/song/ft98y37.html','简弘亦 - 天问【思美人插曲】'),(3,'http://www.kugou.com/song/fteki22.html','张艺兴 - 精忠报国 (Live)'),(4,'http://www.kugou.com/song/ftekj57.html','易烊千玺、王源 - 梦想起航 (Live)'),(5,'http://www.kugou.com/song/ft8fv20.html','黄子韬 - Collateral Love'),(6,'http://www.kugou.com/song/ft7gza8.html','门丽 - 情终人散'),(7,'http://www.kugou.com/song/ft8va44.html','刘心、李炜、谭佑铭、武艺 - 义【毒诫推广曲】'),(8,'http://www.kugou.com/song/ft84u37.html','云飞、李玲玉 - 窗外'),(9,'http://www.kugou.com/song/ft50b67.html','魏晨 - 弟兄【抢红推广曲】'),(10,'http://www.kugou.com/song/5lz0a7e.html','贾征宇、刘明湘 - 漂洋过海来看你(Live)'),(11,'http://www.kugou.com/song/ft0z2ae.html','黑龙 - 盗心贼'),(12,'http://www.kugou.com/song/fsyaze1.html','谭维维 - 昨天涯--献给布宜诺斯艾利斯'),(13,'http://www.kugou.com/song/f99hy3a.html','任宇翔 - 悟空 (试听版)'),(14,'http://www.kugou.com/song/fsygfde.html','单色凌 - 欠你一句抱歉'),(15,'http://www.kugou.com/song/fsxuc67.html','汪苏泷 - 青春白皮书【小情书推广曲】'),(16,'http://www.kugou.com/song/fswic66.html','董贞 - 静女思'),(17,'http://www.kugou.com/song/fsy9gab.html','金志文 - 青春不散'),(18,'http://www.kugou.com/song/fsnl706.html','方圆 - 橘颂【思美人插曲】'),(19,'http://www.kugou.com/song/fswib64.html','陈瑞 - 终于放了手'),(20,'http://www.kugou.com/song/fsy9faf.html','羽·泉 - 青春不散'),(21,'http://www.kugou.com/song/ebmbf65.html','Various Artists - Feded - FeNG (Remix)'),(22,'http://www.kugou.com/song/fsnpebc.html','阿杜 - 烂好人');
/*!40000 ALTER TABLE `kugou_biaoshengbang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kugou_net_hot`
--

DROP TABLE IF EXISTS `kugou_net_hot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kugou_net_hot` (
  `music_id` int(11) NOT NULL AUTO_INCREMENT,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`music_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kugou_net_hot`
--

LOCK TABLES `kugou_net_hot` WRITE;
/*!40000 ALTER TABLE `kugou_net_hot` DISABLE KEYS */;
INSERT INTO `kugou_net_hot` VALUES (1,'http://www.kugou.com/song/1haa0d0.html','金南玲 - 逆流成河'),(2,'http://www.kugou.com/song/b4e7e33.html','庄心妍 - 走着走着就散了'),(3,'http://www.kugou.com/song/fq7q38f.html','唐古 - 你傻不傻'),(4,'http://www.kugou.com/song/fbn2b5.html','庄心妍 - 以后的以后'),(5,'http://www.kugou.com/song/2bl023f.html','神马乐团 - 爱河'),(6,'http://www.kugou.com/song/1sx2z72.html','苏勒亚其其格 - 歌在飞'),(7,'http://www.kugou.com/song/1lgxj91.html','筷子兄弟 - 小苹果'),(8,'http://www.kugou.com/song/fswib64.html','陈瑞 - 终于放了手'),(9,'http://www.kugou.com/song/eo25ec.html','庄心妍 - 爱囚'),(10,'http://www.kugou.com/song/wmu0c5.html','齐晨 - 咱们结婚吧'),(11,'http://www.kugou.com/song/ft7gza8.html','门丽 - 情终人散'),(12,'http://www.kugou.com/song/fswih58.html','Xun - 那年你我'),(13,'http://www.kugou.com/song/fsowoc9.html','浩轩 - 曾经相爱'),(14,'http://www.kugou.com/song/fswip1d.html','龙奔 - 男人太难'),(15,'http://www.kugou.com/song/fswied5.html','张怡诺 - 如果你爱上她'),(16,'http://www.kugou.com/song/gc4c5.html','王童语 - 丫头'),(17,'http://www.kugou.com/song/ft7gv6e.html','赵鑫 - 爱恨不容易'),(18,'http://www.kugou.com/song/702df5.html','许嵩、何曼婷 - 素颜'),(19,'http://www.kugou.com/song/6xj864.html','许嵩 - 断桥残雪'),(20,'http://www.kugou.com/song/coilbf.html','夏天Alex - 不再联系'),(21,'http://www.kugou.com/song/fsoix67.html','周恬熙 - 白露之前'),(22,'http://www.kugou.com/song/fsoid9a.html','于子将 - 双击666');
/*!40000 ALTER TABLE `kugou_net_hot` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kugou_new_song`
--

DROP TABLE IF EXISTS `kugou_new_song`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kugou_new_song` (
  `music_id` int(11) NOT NULL AUTO_INCREMENT,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`music_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kugou_new_song`
--

LOCK TABLES `kugou_new_song` WRITE;
/*!40000 ALTER TABLE `kugou_new_song` DISABLE KEYS */;
INSERT INTO `kugou_new_song` VALUES (1,'http://www.kugou.com/song/ftm9b75.html','李宇春 - 流言'),(2,'http://www.kugou.com/song/fpg3nb6.html','薛之谦 - 我害怕'),(3,'http://www.kugou.com/song/fsxuc67.html','汪苏泷 - 青春白皮书'),(4,'http://www.kugou.com/song/fsyaze1.html','谭维维 - 昨天涯--献给布宜诺斯艾利斯'),(5,'http://www.kugou.com/song/fo6s75e.html','易烊千玺 - 离骚'),(6,'http://www.kugou.com/song/frr5bee.html','张艺兴 - 祈愿'),(7,'http://www.kugou.com/song/ft8va44.html','刘心、李炜、谭佑铭、武艺 - 义'),(8,'http://www.kugou.com/song/fsh5n93.html','何洁、孙郡 - 故意'),(9,'http://www.kugou.com/song/fnys198.html','陈伟霆 - 着迷'),(10,'http://www.kugou.com/song/fs4wda8.html','刘惜君 - 不期而遇'),(11,'http://www.kugou.com/song/fqlqxf2.html','戚薇 - 梦诛缘·春生'),(12,'http://www.kugou.com/song/ft8pyc3.html','徐千雅 - 岁月之影'),(13,'http://www.kugou.com/song/fth0gc8.html','沈梦辰 - 追梦星城'),(14,'http://www.kugou.com/song/ft8xr75.html','徐浩 - 后来呢'),(15,'http://www.kugou.com/song/fr36xd1.html','伦桑 - 神之印记'),(16,'http://www.kugou.com/song/fsnpebc.html','阿杜 - 烂好人'),(17,'http://www.kugou.com/song/fq98m03.html','张涵予、黎明 - 大宝'),(18,'http://www.kugou.com/song/fornh81.html','金志文 - 远走高飞'),(19,'http://www.kugou.com/song/fsm8gad.html','任素汐 - 心恋'),(20,'http://www.kugou.com/song/ft50b67.html','魏晨 - 弟兄'),(21,'http://www.kugou.com/song/ftff8a7.html','蒋蒋、小右 - 你是什么咖'),(22,'http://www.kugou.com/song/fm4m3b1.html','薛之谦 - 暧昧');
/*!40000 ALTER TABLE `kugou_new_song` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kugou_originsong`
--

DROP TABLE IF EXISTS `kugou_originsong`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kugou_originsong` (
  `music_id` int(11) NOT NULL AUTO_INCREMENT,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`music_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kugou_originsong`
--

LOCK TABLES `kugou_originsong` WRITE;
/*!40000 ALTER TABLE `kugou_originsong` DISABLE KEYS */;
INSERT INTO `kugou_originsong` VALUES (1,'http://www.kugou.com/song/e2153ea.html','李玉刚 - 刚好遇见你'),(2,'http://www.kugou.com/song/ewjp8ff.html','赵雷 - 成都 (Live)'),(3,'http://www.kugou.com/song/exul788.html','董贞 - 繁花'),(4,'http://www.kugou.com/song/draum85.html','孙子涵 - 回忆那么伤'),(5,'http://www.kugou.com/song/fornh81.html','金志文 - 远走高飞'),(6,'http://www.kugou.com/song/fp0z6ac.html','戴荃、周华健 - 江湖晚'),(7,'http://www.kugou.com/song/by0qmf3.html','崔子格 - 可念不可说'),(8,'http://www.kugou.com/song/de4v558.html','西瓜JUN - 长生诀'),(9,'http://www.kugou.com/song/fm2gva0.html','霍尊 - 粉墨'),(10,'http://www.kugou.com/song/fs5mwe1.html','徐真真 - 呼吸'),(11,'http://www.kugou.com/song/fqqkmf4.html','金志文 - Hello'),(12,'http://www.kugou.com/song/fczn5f5.html','董贞 - 南风吹'),(13,'http://www.kugou.com/song/epqgnac.html','花粥 - 遥不可及的你'),(14,'http://www.kugou.com/song/fsxxwa1.html','万晓利 - 你,来替我做个梦'),(15,'http://www.kugou.com/song/fn96qb1.html','萧忆情Alex、玄觞、祈inory、洛少爷 - 萌动西域'),(16,'http://www.kugou.com/song/eqznvaf.html','周品 - 不负'),(17,'http://www.kugou.com/song/fk7ira7.html','叶洛洛 - 梨花凉'),(18,'http://www.kugou.com/song/dcy4205.html','徐菲 - 树叶的光'),(19,'http://www.kugou.com/song/d9n1d38.html','银临 - 不老梦'),(20,'http://www.kugou.com/song/fmdtp85.html','冰块先生、郭美孜 - Hey Jude'),(21,'http://www.kugou.com/song/fftyg69.html','萧忆情Alex - 拜无忧'),(22,'http://www.kugou.com/song/at0fcc0.html','丢火车乐队 - 晚安');
/*!40000 ALTER TABLE `kugou_originsong` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kugou_top500`
--

DROP TABLE IF EXISTS `kugou_top500`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kugou_top500` (
  `music_id` int(11) NOT NULL AUTO_INCREMENT,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`music_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kugou_top500`
--

LOCK TABLES `kugou_top500` WRITE;
/*!40000 ALTER TABLE `kugou_top500` DISABLE KEYS */;
INSERT INTO `kugou_top500` VALUES (1,'http://www.kugou.com/song/eqi2pe8.html','杨宗纬、张碧晨 - 凉凉'),(2,'http://www.kugou.com/song/btqik70.html','赵雷 - 成都'),(3,'http://www.kugou.com/song/e2153ea.html','李玉刚 - 刚好遇见你'),(4,'http://www.kugou.com/song/2bl023f.html','神马乐团 - 爱河'),(5,'http://www.kugou.com/song/fm4m3b1.html','薛之谦 - 暧昧'),(6,'http://www.kugou.com/song/fphka77.html','夏天Alex - 爱河'),(7,'http://www.kugou.com/song/94ikj9f.html','薛之谦 - 演员'),(8,'http://www.kugou.com/song/d5c5m23.html','周杰伦 - 告白气球'),(9,'http://www.kugou.com/song/9qe65f3.html','陈雅森 - 温柔乡'),(10,'http://www.kugou.com/song/br17856.html','Alan Walker - Faded'),(11,'http://www.kugou.com/song/6aw3xda.html','萧全 - 社会摇'),(12,'http://www.kugou.com/song/as81813.html','鹿晗 - 致爱 Your Song'),(13,'http://www.kugou.com/song/dzezz1e.html','王冕 - 勉为其难'),(14,'http://www.kugou.com/song/b29yqdd.html','李晓杰 - 把酒倒满'),(15,'http://www.kugou.com/song/fpg3nb6.html','薛之谦 - 我害怕'),(16,'http://www.kugou.com/song/1haa0d0.html','金南玲 - 逆流成河'),(17,'http://www.kugou.com/song/eydc64.html','许云上 - 爱河'),(18,'http://www.kugou.com/song/1sx2z72.html','苏勒亚其其格 - 歌在飞'),(19,'http://www.kugou.com/song/l3d2b1.html','梦然 - 没有你陪伴真的好孤单'),(20,'http://www.kugou.com/song/feku0b6.html','王丽坤、朱亚文 - 漂洋过海来看你'),(21,'http://www.kugou.com/song/erfb00b.html','张杰 - 三生三世'),(22,'http://www.kugou.com/song/ccm7q60.html','葛林 - 林中鸟');
/*!40000 ALTER TABLE `kugou_top500` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-05-07 15:27:26
