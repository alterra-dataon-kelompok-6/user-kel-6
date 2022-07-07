-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: kelompok_6
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_role_id` bigint DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `name` varchar(200) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(200) NOT NULL,
  `address` varchar(200) DEFAULT NULL,
  `password` varchar(200) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `token` varchar(700) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_UserRoleId` (`user_role_id`),
  CONSTRAINT `FK_UserRoleId` FOREIGN KEY (`user_role_id`) REFERENCES `user_roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,1,'azkamuhamco','Muhammad Azka Ramadhan','085155188431','mazkar.indonesia@gmail.com','Yogyakarta Nyaman','1234567','2022-06-30 14:54:35','2022-06-30 23:21:28',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY2MzQ4ODcsInVzZXJJZCI6MX0.pmPbEi4v_t32P7-w0upP-hmV9umwKrtMGAd9wDZCAw0'),(2,1,'rayoda','Rayoda Adiwarna Rahmazka','0812345678','rayoda@azka.com','Depok Jawa Barat','7654321','2022-06-30 23:33:49','2022-06-30 23:42:26',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY2MzYxNDUsInVzZXJJZCI6Mn0.Aj6aLAyex3xq3bU1ypw8qVr8OpCPZROwe_fDbX6ZZNE'),(6,2,'karim.muhamr','Karim Muhammad','0812345678','karim@azka.com','Notoprajan, Yogyakarta','karim','2022-07-06 15:40:50','2022-07-06 15:47:18',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTcxMjYwMzcsInVzZXJJZCI6Nn0.aOlhqy4B917_et1LyCzeYTWWhZCa1d-FfJNAcBadTg8'),(7,2,'rahmasari.azka','Rahmawati Purnamasari','087878781212','rahmasari@azka.com','Madiun','7654321','2022-07-06 22:48:32','2022-07-06 22:49:51',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTcxNTEzOTAsInVzZXJJZCI6N30.xWhekiVsNRfowzfWWaOcoBR65hxrVjfGK9oqKLitOR4'),(8,2,'pittra','Pittra Amrullah','081878781212','pittra@azka.com','Notoprajan, Yogyakarta','pittraburhan','2022-07-06 22:56:42','2022-07-06 22:57:09',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTcxNTE4MjksInVzZXJJZCI6OH0.HAsmCdpys5TzLNyAh9MjWCqHp-PUN-1f2R_XucgJVbE'),(9,2,'sahal','Sahal Mahfudh','081878781212','sahal@mahfudh.com','Alterra Malang','lolosdataon','2022-07-07 04:48:28','2022-07-07 05:48:16',NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTcxNzUyNjksInVzZXJJZCI6OX0.rWi2HdmDunXh38M6JLV7sUlZOwa4_-ykufeQ7DYcYCI'),(10,2,'sahal','Sahal Mahfudh','081878781212','sahal@azka.com','Alterra Malang','lolosdataon','2022-07-07 05:38:27','2022-07-07 05:38:27',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-07 13:29:59
