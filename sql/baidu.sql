-- MySQL dump 10.16  Distrib 10.1.22-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: baidu
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
-- Table structure for table `baidu_hotsong`
--

DROP TABLE IF EXISTS `baidu_hotsong`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `baidu_hotsong` (
  `music_id` int(11) DEFAULT NULL,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  `music_artist` varchar(255) DEFAULT NULL,
  `music_artist_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baidu_hotsong`
--

LOCK TABLES `baidu_hotsong` WRITE;
/*!40000 ALTER TABLE `baidu_hotsong` DISABLE KEYS */;
INSERT INTO `baidu_hotsong` VALUES (1,'http://music.baidu.com/song/540560826','我害怕','薛之谦','http://music.baidu.com/artist/2517'),(2,'http://music.baidu.com/song/540175998','暧昧','薛之谦','http://music.baidu.com/artist/2517'),(3,'http://music.baidu.com/song/540489526','春风十里不如你','李健','http://music.baidu.com/artist/1383'),(4,'http://music.baidu.com/song/540430432','我不能忘记你','林忆莲','http://music.baidu.com/artist/1133'),(5,'http://music.baidu.com/song/540505067','飞天','云朵','http://music.baidu.com/artist/177498'),(6,'http://music.baidu.com/song/540565342','Lisa(Remix)','丁于','http://music.baidu.com/artist/1672'),(7,'http://music.baidu.com/song/540159170','失落的缘','谭维维','http://music.baidu.com/artist/1062'),(8,'http://music.baidu.com/song/539843843','动物世界','薛之谦','http://music.baidu.com/artist/2517'),(9,'http://music.baidu.com/song/540076960','爱到五十年以后','刘明辉','http://music.baidu.com/artist/1389007'),(10,'http://music.baidu.com/song/540489965','瞄着你就爱','凤凰传奇','http://music.baidu.com/artist/1490');
/*!40000 ALTER TABLE `baidu_hotsong` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `baidu_netsong`
--

DROP TABLE IF EXISTS `baidu_netsong`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `baidu_netsong` (
  `music_id` int(11) DEFAULT NULL,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  `music_artist` varchar(255) DEFAULT NULL,
  `music_artist_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baidu_netsong`
--

LOCK TABLES `baidu_netsong` WRITE;
/*!40000 ALTER TABLE `baidu_netsong` DISABLE KEYS */;
INSERT INTO `baidu_netsong` VALUES (1,'http://music.baidu.com/song/540560826','我害怕','薛之谦','http://music.baidu.com/artist/2517'),(2,'http://music.baidu.com/song/540175998','暧昧','薛之谦','http://music.baidu.com/artist/2517'),(3,'http://music.baidu.com/song/540489526','春风十里不如你','李健','http://music.baidu.com/artist/1383'),(4,'http://music.baidu.com/song/540430432','我不能忘记你','林忆莲','http://music.baidu.com/artist/1133'),(5,'http://music.baidu.com/song/540505067','飞天','云朵','http://music.baidu.com/artist/177498'),(6,'http://music.baidu.com/song/540565342','Lisa(Remix)','丁于','http://music.baidu.com/artist/1672'),(7,'http://music.baidu.com/song/540159170','失落的缘','谭维维','http://music.baidu.com/artist/1062'),(8,'http://music.baidu.com/song/539843843','动物世界','薛之谦','http://music.baidu.com/artist/2517'),(9,'http://music.baidu.com/song/540076960','爱到五十年以后','刘明辉','http://music.baidu.com/artist/1389007'),(10,'http://music.baidu.com/song/540489965','瞄着你就爱','凤凰传奇','http://music.baidu.com/artist/1490');
/*!40000 ALTER TABLE `baidu_netsong` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `baidu_newsong`
--

DROP TABLE IF EXISTS `baidu_newsong`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `baidu_newsong` (
  `music_id` int(11) DEFAULT NULL,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  `music_artist` varchar(255) DEFAULT NULL,
  `music_artist_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baidu_newsong`
--

LOCK TABLES `baidu_newsong` WRITE;
/*!40000 ALTER TABLE `baidu_newsong` DISABLE KEYS */;
INSERT INTO `baidu_newsong` VALUES (1,'http://music.baidu.com/song/540560826','我害怕','薛之谦','http://music.baidu.com/artist/2517'),(2,'http://music.baidu.com/song/540175998','暧昧','薛之谦','http://music.baidu.com/artist/2517'),(3,'http://music.baidu.com/song/540489526','春风十里不如你','李健','http://music.baidu.com/artist/1383'),(4,'http://music.baidu.com/song/540430432','我不能忘记你','林忆莲','http://music.baidu.com/artist/1133'),(5,'http://music.baidu.com/song/540505067','飞天','云朵','http://music.baidu.com/artist/177498'),(6,'http://music.baidu.com/song/540565342','Lisa(Remix)','丁于','http://music.baidu.com/artist/1672'),(7,'http://music.baidu.com/song/540159170','失落的缘','谭维维','http://music.baidu.com/artist/1062'),(8,'http://music.baidu.com/song/539843843','动物世界','薛之谦','http://music.baidu.com/artist/2517'),(9,'http://music.baidu.com/song/540076960','爱到五十年以后','刘明辉','http://music.baidu.com/artist/1389007'),(10,'http://music.baidu.com/song/540489965','瞄着你就爱','凤凰传奇','http://music.baidu.com/artist/1490');
/*!40000 ALTER TABLE `baidu_newsong` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `baidu_originsong`
--

DROP TABLE IF EXISTS `baidu_originsong`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `baidu_originsong` (
  `music_id` int(11) DEFAULT NULL,
  `music_url` varchar(255) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  `music_artist` varchar(255) DEFAULT NULL,
  `music_artist_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baidu_originsong`
--

LOCK TABLES `baidu_originsong` WRITE;
/*!40000 ALTER TABLE `baidu_originsong` DISABLE KEYS */;
INSERT INTO `baidu_originsong` VALUES (1,'http://music.baidu.com/song/540560826','我害怕','薛之谦','http://music.baidu.com/artist/2517'),(2,'http://music.baidu.com/song/540175998','暧昧','薛之谦','http://music.baidu.com/artist/2517'),(3,'http://music.baidu.com/song/540489526','春风十里不如你','李健','http://music.baidu.com/artist/1383'),(4,'http://music.baidu.com/song/540430432','我不能忘记你','林忆莲','http://music.baidu.com/artist/1133'),(5,'http://music.baidu.com/song/540505067','飞天','云朵','http://music.baidu.com/artist/177498'),(6,'http://music.baidu.com/song/540565342','Lisa(Remix)','丁于','http://music.baidu.com/artist/1672'),(7,'http://music.baidu.com/song/540159170','失落的缘','谭维维','http://music.baidu.com/artist/1062'),(8,'http://music.baidu.com/song/539843843','动物世界','薛之谦','http://music.baidu.com/artist/2517'),(9,'http://music.baidu.com/song/540076960','爱到五十年以后','刘明辉','http://music.baidu.com/artist/1389007'),(10,'http://music.baidu.com/song/540489965','瞄着你就爱','凤凰传奇','http://music.baidu.com/artist/1490');
/*!40000 ALTER TABLE `baidu_originsong` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `music`
--

DROP TABLE IF EXISTS `music`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `music` (
  `music_id` int(11) DEFAULT NULL,
  `music_name` varchar(255) DEFAULT NULL,
  `music_artist` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `music`
--

LOCK TABLES `music` WRITE;
/*!40000 ALTER TABLE `music` DISABLE KEYS */;
INSERT INTO `music` VALUES (1,'我害怕','薛之谦'),(2,'暧昧','薛之谦'),(3,'春风十里不如你','李健'),(4,'我不能忘记你','林忆莲'),(5,'飞天','云朵'),(6,'Lisa(Remix)','丁于'),(7,'失落的缘','谭维维'),(8,'动物世界','薛之谦'),(9,'爱到五十年以后','刘明辉'),(10,'瞄着你就爱','凤凰传奇'),(1,'刚好遇见你','李玉刚'),(2,'成都','赵雷'),(3,'暧昧','薛之谦'),(4,'告白气球','周杰伦'),(5,'修今生','梁佳玉'),(6,'演员','薛之谦'),(7,'因为遇见你','阿悄,朱元冰'),(8,'一生为你感动','祁隆'),(9,'寂寞的人伤心的歌','龙梅子,杨海彪'),(10,'怎样遇见你','孙露'),(1,'谁说','硬摇滚'),(2,'在四季里等待','独立民谣'),(3,'王子与河童','独立流行'),(4,'人海','国语流行'),(5,'Another Night','爵士流行'),(6,'花大侠','独立流行'),(7,'看你交的什么朋友','国语流行'),(8,'深海里的一条鱼','国语流行'),(9,'Run Run Run','放克'),(10,'火花','国语流行'),(1,'独角戏','许茹芸'),(2,'一路上有你','张学友'),(3,'倩女幽魂','张国荣'),(4,'相思风雨中','张学友,汤宝如'),(5,'光辉岁月','Beyond'),(6,'无情的雨无情的你','齐秦'),(7,'我只在乎你','邓丽君'),(8,'容易受伤的女人','王菲'),(9,'偏偏喜欢你','陈百强'),(10,'灰姑娘','郑钧'),(1,'因为遇见你','阿悄,朱元冰'),(2,'怎样遇见你','孙露'),(3,'一生为你感动','祁隆'),(4,'逆流成河','宇桐非'),(5,'就是让你美','龙梅子'),(6,'今生遇见你','祁隆,任妙音'),(7,'寂寞的人伤心的歌','龙梅子,杨海彪'),(8,'瞄着你就爱','凤凰传奇'),(9,'大叔不卖我香蕉','龙梅子,老猫'),(10,'下完这场雨','后弦'),(1,'过境','愚青'),(2,'相思入画','百花楼原创音乐团队'),(3,'雪女之雪魇','小爱的妈'),(4,'秋城','丢火车乐队'),(5,'洗脑的玩具','蔡丰'),(6,'【择天记】长生愿（唱：西瓜Kune）','小玖州'),(7,'琵琶行','奇然'),(8,'君是山','戴荃▪悟空'),(9,'恰遇斜阳照深林','长歌红影乱'),(10,'听无可说','河图'),(1,'我害怕','薛之谦'),(2,'暧昧','薛之谦'),(3,'春风十里不如你','李健'),(4,'我不能忘记你','林忆莲'),(5,'飞天','云朵'),(6,'Lisa(Remix)','丁于'),(7,'失落的缘','谭维维'),(8,'动物世界','薛之谦'),(9,'爱到五十年以后','刘明辉'),(10,'瞄着你就爱','凤凰传奇'),(1,'刚好遇见你','李玉刚'),(2,'成都','赵雷'),(3,'暧昧','薛之谦'),(4,'告白气球','周杰伦'),(5,'修今生','梁佳玉'),(6,'演员','薛之谦'),(7,'因为遇见你','阿悄,朱元冰'),(8,'一生为你感动','祁隆'),(9,'寂寞的人伤心的歌','龙梅子,杨海彪'),(10,'怎样遇见你','孙露'),(1,'谁说','硬摇滚'),(2,'在四季里等待','独立民谣'),(3,'王子与河童','独立流行'),(4,'人海','国语流行'),(5,'Another Night','爵士流行'),(6,'花大侠','独立流行'),(7,'看你交的什么朋友','国语流行'),(8,'深海里的一条鱼','国语流行'),(9,'Run Run Run','放克'),(10,'火花','国语流行'),(1,'独角戏','许茹芸'),(2,'一路上有你','张学友'),(3,'倩女幽魂','张国荣'),(4,'相思风雨中','张学友,汤宝如'),(5,'光辉岁月','Beyond'),(6,'无情的雨无情的你','齐秦'),(7,'我只在乎你','邓丽君'),(8,'容易受伤的女人','王菲'),(9,'偏偏喜欢你','陈百强'),(10,'灰姑娘','郑钧'),(1,'因为遇见你','阿悄,朱元冰'),(2,'怎样遇见你','孙露'),(3,'一生为你感动','祁隆'),(4,'逆流成河','宇桐非'),(5,'就是让你美','龙梅子'),(6,'今生遇见你','祁隆,任妙音'),(7,'寂寞的人伤心的歌','龙梅子,杨海彪'),(8,'瞄着你就爱','凤凰传奇'),(9,'大叔不卖我香蕉','龙梅子,老猫'),(10,'下完这场雨','后弦'),(1,'过境','愚青'),(2,'相思入画','百花楼原创音乐团队'),(3,'雪女之雪魇','小爱的妈'),(4,'秋城','丢火车乐队'),(5,'洗脑的玩具','蔡丰'),(6,'【择天记】长生愿（唱：西瓜Kune）','小玖州'),(7,'琵琶行','奇然'),(8,'君是山','戴荃▪悟空'),(9,'恰遇斜阳照深林','长歌红影乱'),(10,'听无可说','河图'),(1,'/song/540560826','我害怕'),(2,'/song/540175998','暧昧'),(3,'/song/540489526','春风十里不如你'),(4,'/song/540430432','我不能忘记你'),(5,'/song/540505067','飞天'),(6,'/song/540565342','Lisa(Remix)'),(7,'/song/540159170','失落的缘'),(8,'/song/539843843','动物世界'),(9,'/song/540076960','爱到五十年以后'),(10,'/song/540489965','瞄着你就爱'),(1,'/song/276867440','刚好遇见你'),(2,'/song/274841326','成都'),(3,'/song/540175998','暧昧'),(4,'/song/266322598','告白气球'),(5,'/song/277058670','修今生'),(6,'/song/242078437','演员'),(7,'/song/537792416','因为遇见你'),(8,'/song/277389316','一生为你感动'),(9,'/song/265046969','寂寞的人伤心的歌'),(10,'/song/274334244','怎样遇见你'),(1,'谁说','硬摇滚'),(2,'在四季里等待','独立民谣'),(3,'王子与河童','独立流行'),(4,'人海','国语流行'),(5,'Another Night','爵士流行'),(6,'花大侠','独立流行'),(7,'看你交的什么朋友','国语流行'),(8,'深海里的一条鱼','国语流行'),(9,'Run Run Run','放克'),(10,'火花','国语流行'),(1,'/song/490468','独角戏'),(2,'/song/620023','一路上有你'),(3,'/song/13125209','倩女幽魂'),(4,'/song/7317902','相思风雨中'),(5,'/song/7338336','光辉岁月'),(6,'/song/1183139','无情的雨无情的你'),(7,'/song/290008','我只在乎你'),(8,'/song/307171','容易受伤的女人'),(9,'/song/7320512','偏偏喜欢你'),(10,'/song/276766','灰姑娘'),(1,'/song/537792416','因为遇见你'),(2,'/song/274334244','怎样遇见你'),(3,'/song/277389316','一生为你感动'),(4,'/song/533308233','逆流成河'),(5,'/song/269266591','就是让你美'),(6,'/song/271896483','今生遇见你'),(7,'/song/265046969','寂寞的人伤心的歌'),(8,'/song/540489965','瞄着你就爱'),(9,'/song/270882096','大叔不卖我香蕉'),(10,'/song/272952711','下完这场雨'),(1,'/song/74200488/?pst=shouyeTop','过境'),(2,'/song/74115535/?pst=shouyeTop','相思入画'),(3,'/song/74205450/?pst=shouyeTop','雪女之雪魇'),(4,'/song/74205302/?pst=shouyeTop','秋城'),(5,'/song/74205354/?pst=shouyeTop','洗脑的玩具'),(6,'/song/74204734/?pst=shouyeTop','【择天记】长生愿（唱：西瓜Kune）'),(7,'/song/74207004/?pst=shouyeTop','琵琶行'),(8,'/song/74204864/?pst=shouyeTop','君是山'),(9,'/song/74204753/?pst=shouyeTop','恰遇斜阳照深林'),(10,'/song/74205754/?pst=shouyeTop','听无可说'),(1,'/song/540560826','我害怕'),(2,'/song/540175998','暧昧'),(3,'/song/540489526','春风十里不如你'),(4,'/song/540430432','我不能忘记你'),(5,'/song/540505067','飞天'),(6,'/song/540565342','Lisa(Remix)'),(7,'/song/540159170','失落的缘'),(8,'/song/539843843','动物世界'),(9,'/song/540076960','爱到五十年以后'),(10,'/song/540489965','瞄着你就爱'),(1,'/song/276867440','刚好遇见你'),(2,'/song/274841326','成都'),(3,'/song/540175998','暧昧'),(4,'/song/266322598','告白气球'),(5,'/song/277058670','修今生'),(6,'/song/242078437','演员'),(7,'/song/537792416','因为遇见你'),(8,'/song/277389316','一生为你感动'),(9,'/song/265046969','寂寞的人伤心的歌'),(10,'/song/274334244','怎样遇见你'),(1,'谁说','硬摇滚'),(2,'在四季里等待','独立民谣'),(3,'王子与河童','独立流行'),(4,'人海','国语流行'),(5,'Another Night','爵士流行'),(6,'花大侠','独立流行'),(7,'看你交的什么朋友','国语流行'),(8,'深海里的一条鱼','国语流行'),(9,'Run Run Run','放克'),(10,'火花','国语流行'),(1,'/song/490468','独角戏'),(2,'/song/620023','一路上有你'),(3,'/song/13125209','倩女幽魂'),(4,'/song/7317902','相思风雨中'),(5,'/song/7338336','光辉岁月'),(6,'/song/1183139','无情的雨无情的你'),(7,'/song/290008','我只在乎你'),(8,'/song/307171','容易受伤的女人'),(9,'/song/7320512','偏偏喜欢你'),(10,'/song/276766','灰姑娘'),(1,'/song/537792416','因为遇见你'),(2,'/song/274334244','怎样遇见你'),(3,'/song/277389316','一生为你感动'),(4,'/song/533308233','逆流成河'),(5,'/song/269266591','就是让你美'),(6,'/song/271896483','今生遇见你'),(7,'/song/265046969','寂寞的人伤心的歌'),(8,'/song/540489965','瞄着你就爱'),(9,'/song/270882096','大叔不卖我香蕉'),(10,'/song/272952711','下完这场雨'),(1,'/song/74200488/?pst=shouyeTop','过境'),(2,'/song/74115535/?pst=shouyeTop','相思入画'),(3,'/song/74205450/?pst=shouyeTop','雪女之雪魇'),(4,'/song/74205302/?pst=shouyeTop','秋城'),(5,'/song/74205354/?pst=shouyeTop','洗脑的玩具'),(6,'/song/74204734/?pst=shouyeTop','【择天记】长生愿（唱：西瓜Kune）'),(7,'/song/74207004/?pst=shouyeTop','琵琶行'),(8,'/song/74204864/?pst=shouyeTop','君是山'),(9,'/song/74204753/?pst=shouyeTop','恰遇斜阳照深林'),(10,'/song/74205754/?pst=shouyeTop','听无可说'),(1,'/song/540560826','我害怕'),(2,'/song/540175998','暧昧'),(3,'/song/540489526','春风十里不如你'),(4,'/song/540430432','我不能忘记你'),(5,'/song/540505067','飞天'),(6,'/song/540565342','Lisa(Remix)'),(7,'/song/540159170','失落的缘'),(8,'/song/539843843','动物世界'),(9,'/song/540076960','爱到五十年以后'),(10,'/song/540489965','瞄着你就爱');
/*!40000 ALTER TABLE `music` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-05-07 15:27:45
