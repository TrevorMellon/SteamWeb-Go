--
-- Current Database: `coralinesteam`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `coralinesteam` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

USE `coralinesteam`;

--
-- Table structure for table `friends`
--

DROP TABLE IF EXISTS `friends`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `friends` (
  `steamid` bigint(64) NOT NULL,
  `friendid` bigint(64) NOT NULL,
  `friendssince` datetime(6) NOT NULL,
  `dataadded` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`steamid`,`friendid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;



--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `steamid` bigint(64) NOT NULL,
  `personaname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `realname` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `personanameblob` blob,
  `realnameblob` blob,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `summaryparsed` datetime(6) DEFAULT NULL,
  `dateadded` datetime(6) DEFAULT NULL,
  `avsmall` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avmedium` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avlarge` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profiletype` smallint(3) DEFAULT '0',
  `status` smallint(3) DEFAULT '0',
  PRIMARY KEY (`steamid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `usergames`
--

DROP TABLE IF EXISTS `usergames`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usergames` (
  `steamid` bigint(64) NOT NULL,
  `appid` bigint(64) NOT NULL,
  `playtime` bigint(64) NOT NULL DEFAULT '0',
  PRIMARY KEY (`steamid`,`appid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;


