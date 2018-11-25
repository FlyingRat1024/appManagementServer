-- MySQL dump 10.13  Distrib 5.7.16, for Linux (x86_64)
--
-- Host: localhost    Database: app_server
-- ------------------------------------------------------
-- Server version	5.7.16-0ubuntu2

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
-- Table structure for table `apply_material`
--

DROP TABLE IF EXISTS `apply_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `apply_material` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table_id` int(11) DEFAULT NULL,
  `material_id` int(11) DEFAULT NULL,
  `num` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `table_id` (`table_id`),
  KEY `material_id` (`material_id`),
  CONSTRAINT `apply_material_ibfk_1` FOREIGN KEY (`table_id`) REFERENCES `material_apply_table` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `apply_material_ibfk_2` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apply_material`
--

LOCK TABLES `apply_material` WRITE;
/*!40000 ALTER TABLE `apply_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `apply_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `in_material`
--

DROP TABLE IF EXISTS `in_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `in_material` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `in_id` int(11) DEFAULT NULL,
  `material_id` int(11) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `in_id` (`in_id`),
  KEY `material_id` (`material_id`),
  CONSTRAINT `in_material_ibfk_1` FOREIGN KEY (`in_id`) REFERENCES `warehouse_in` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `in_material_ibfk_2` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `in_material`
--

LOCK TABLES `in_material` WRITE;
/*!40000 ALTER TABLE `in_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `in_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `material`
--

DROP TABLE IF EXISTS `material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `material` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `unit` varchar(8) DEFAULT NULL,
  `provider` varchar(255) DEFAULT NULL,
  `description` varchar(155) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `material`
--

LOCK TABLES `material` WRITE;
/*!40000 ALTER TABLE `material` DISABLE KEYS */;
/*!40000 ALTER TABLE `material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `material_apply_table`
--

DROP TABLE IF EXISTS `material_apply_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `material_apply_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `create_time` datetime NOT NULL,
  `verifier` int(11) DEFAULT NULL,
  `verify_time` datetime DEFAULT NULL,
  `verify` int(11) DEFAULT NULL COMMENT '是否审核通过0未审核，1通过，-1未通过',
  `verify_comment` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `verifier` (`verifier`),
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `verifier` FOREIGN KEY (`verifier`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `material_apply_table`
--

LOCK TABLES `material_apply_table` WRITE;
/*!40000 ALTER TABLE `material_apply_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `material_apply_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `material_receive_table`
--

DROP TABLE IF EXISTS `material_receive_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `material_receive_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `receiver` int(11) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `verifier` int(11) DEFAULT NULL,
  `verify` int(11) DEFAULT NULL COMMENT '是否审核通过0未审核，1通过，-1未通过',
  `verify_time` datetime DEFAULT NULL,
  `verify_comment` varchar(255) DEFAULT NULL,
  `checker` int(11) DEFAULT NULL,
  `check_time` datetime DEFAULT NULL,
  `check` tinyint(4) DEFAULT NULL COMMENT '1:检查过了\n2：未检查',
  `back_user` int(11) DEFAULT NULL,
  `back_time` datetime DEFAULT NULL,
  `back` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `receiver` (`receiver`),
  KEY `checker` (`checker`),
  KEY `verifier` (`verifier`),
  KEY `back_user` (`back_user`),
  CONSTRAINT `material_receive_table_ibfk_1` FOREIGN KEY (`receiver`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `material_receive_table_ibfk_2` FOREIGN KEY (`checker`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `material_receive_table_ibfk_3` FOREIGN KEY (`verifier`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `material_receive_table_ibfk_4` FOREIGN KEY (`back_user`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `material_receive_table`
--

LOCK TABLES `material_receive_table` WRITE;
/*!40000 ALTER TABLE `material_receive_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `material_receive_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(500) DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `message_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `out_material`
--

DROP TABLE IF EXISTS `out_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `out_material` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `out_id` int(11) DEFAULT NULL,
  `material_id` int(11) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `out_id` (`out_id`),
  KEY `material_id` (`material_id`),
  CONSTRAINT `out_material_ibfk_1` FOREIGN KEY (`out_id`) REFERENCES `warehouse_out` (`id`),
  CONSTRAINT `out_material_ibfk_2` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `out_material`
--

LOCK TABLES `out_material` WRITE;
/*!40000 ALTER TABLE `out_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `out_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receive_material`
--

DROP TABLE IF EXISTS `receive_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `receive_material` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table_id` int(11) NOT NULL,
  `material_id` int(11) NOT NULL,
  `receive_num` int(11) NOT NULL,
  `back_num` int(11) DEFAULT NULL,
  `check_num` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `table_id` (`table_id`),
  KEY `material_id` (`material_id`),
  CONSTRAINT `receive_material_ibfk_1` FOREIGN KEY (`table_id`) REFERENCES `material_receive_table` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `receive_material_ibfk_2` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receive_material`
--

LOCK TABLES `receive_material` WRITE;
/*!40000 ALTER TABLE `receive_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `receive_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `employee_id` varchar(40) NOT NULL,
  `employee_name` varchar(20) NOT NULL,
  `sex` enum('male','female') DEFAULT NULL,
  `position` varchar(45) NOT NULL,
  `password` varchar(255) NOT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'123456','admin','male','root','123456','2018-11-12 00:01:32','2018-11-12 00:01:32');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse`
--

DROP TABLE IF EXISTS `warehouse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warehouse` (
  `material_id` int(11) NOT NULL,
  `num` int(11) DEFAULT NULL,
  `id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `material_id` (`material_id`),
  CONSTRAINT `material_id` FOREIGN KEY (`material_id`) REFERENCES `material` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse`
--

LOCK TABLES `warehouse` WRITE;
/*!40000 ALTER TABLE `warehouse` DISABLE KEYS */;
/*!40000 ALTER TABLE `warehouse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse_in`
--

DROP TABLE IF EXISTS `warehouse_in`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warehouse_in` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `writer` int(11) NOT NULL,
  `reissue` tinyint(4) DEFAULT '0' COMMENT '判断是否为补办的单据',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `writer` (`writer`),
  CONSTRAINT `warehouse_in_ibfk_1` FOREIGN KEY (`writer`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse_in`
--

LOCK TABLES `warehouse_in` WRITE;
/*!40000 ALTER TABLE `warehouse_in` DISABLE KEYS */;
/*!40000 ALTER TABLE `warehouse_in` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse_out`
--

DROP TABLE IF EXISTS `warehouse_out`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warehouse_out` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `writer` int(11) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `verifier` int(11) NOT NULL,
  `verify_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `writer` (`writer`),
  KEY `verifier` (`verifier`),
  CONSTRAINT `warehouse_out_ibfk_1` FOREIGN KEY (`writer`) REFERENCES `user` (`id`),
  CONSTRAINT `warehouse_out_ibfk_2` FOREIGN KEY (`verifier`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse_out`
--

LOCK TABLES `warehouse_out` WRITE;
/*!40000 ALTER TABLE `warehouse_out` DISABLE KEYS */;
/*!40000 ALTER TABLE `warehouse_out` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-11-25 17:36:14
