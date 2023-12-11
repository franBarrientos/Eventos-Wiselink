-- MySQL dump 10.13  Distrib 8.0.35, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: wiselink-events
-- ------------------------------------------------------
-- Server version	8.2.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(40) NOT NULL,
  `short_description` varchar(80) NOT NULL,
  `long_description` varchar(200) NOT NULL,
  `date` datetime(3) DEFAULT NULL,
  `organizer_id` bigint DEFAULT NULL,
  `place_id` bigint DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_events_organizer` (`organizer_id`),
  KEY `fk_events_place` (`place_id`),
  CONSTRAINT `fk_events_organizer` FOREIGN KEY (`organizer_id`) REFERENCES `organizers` (`id`),
  CONSTRAINT `fk_events_place` FOREIGN KEY (`place_id`) REFERENCES `places` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` (`id`, `title`, `short_description`, `long_description`, `date`, `organizer_id`, `place_id`, `state`) VALUES (2,'Introduction To Machine Learning','Learn the basics of ML concepts and algorithms','Join us for an introductory workshop on machine learning. Explore fundamental concepts and algorithms in a hands-on environment.','2023-12-22 17:17:00.000',2,3,1),(7,'AI Ethics Symposium','Discussing ethical considerations in AI development','Engage in thoughtful discussions about the ethical implications of artificial intelligence. Explore real-world case studies and contribute to the conversation','2023-12-22 17:17:05.097',7,8,1),(8,'Blockchain Basics Workshop','Understanding the fundamentals of blockchain technology.','Dive into the world of blockchain with this workshop. Learn about decentralized ledgers, smart contracts, and their applications.','2023-12-20 17:17:05.097',8,9,1),(9,'Python for Data Science','Harness the power of Python for data analysis.','Explore the use of Python in data science. Covering data manipulation, visualization, and introductory machine learning concepts','2023-12-06 17:17:05.097',9,10,1),(10,'Cybersecurity Essentials','Equip yourself with essential cybersecurity knowledge.','pellentesque viverra pede ac diam cras pellentesque volutpat dui maecenas tristique est et tempus semper est quam pharetra','2023-12-22 17:17:00.000',10,11,1),(11,'Deep Learning Workshop','Delving into advanced deep learning techniques','An in-depth workshop on deep learning. Covering neural networks, convolutional networks, and recurrent networks','2023-12-22 17:17:05.097',11,12,1),(12,'Cloud Computing Seminar','Navigating the landscape of cloud technologies','Join us for a seminar on cloud computing. Explore various cloud services, deployment models, and best practices','2023-12-22 17:17:05.097',12,13,1),(26,'IoT Innovation Summit','Exploring the future of the Internet of Things','Participate in discussions and hands-on sessions on the latest trends and innovations in the Internet of Things ecosystem.','2023-12-03 17:17:05.097',8,9,1),(27,'ReactJS Masterclass','architect bleeding-edge conver','A comprehensive masterclass on ReactJS. Learn to build dynamic and efficient user interfaces with this popular JavaScript library.','2023-12-22 17:17:05.097',9,10,1),(28,'AI in Healthcare Symposium','Exploring the impact of AI in the healthcare industry.','Engage with experts and practitioners to discuss the transformative role of artificial intelligence in healthcare. Explore applications, challenges, and future trends.','2023-12-22 17:17:05.097',10,11,1),(29,'DevOps Bootcamp','Bridging the gap between development and operations','Immerse yourself in the principles and practices of DevOps. Learn to streamline collaboration between development and operations teams','2023-12-22 17:17:05.097',11,12,1),(30,'Natural Language Processing Workshop','Unleashing the power of NLP in modern applications.','A hands-on workshop on natural language processing. Explore techniques for text analysis, sentiment analysis, and language understanding.','2023-12-22 17:17:05.097',12,13,0),(31,'Quantum Computing Exploration','Journey into the world of quantum computing.','Embark on a fascinating exploration of quantum computing. Understand the principles, algorithms, and potential applications in this emerging field.','2023-12-22 17:17:00.000',12,11,1);
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizers`
--

DROP TABLE IF EXISTS `organizers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `personal_data_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_organizers_personal_data` (`personal_data_id`),
  CONSTRAINT `fk_organizers_personal_data` FOREIGN KEY (`personal_data_id`) REFERENCES `personal_data` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizers`
--

LOCK TABLES `organizers` WRITE;
/*!40000 ALTER TABLE `organizers` DISABLE KEYS */;
INSERT INTO `organizers` (`id`, `personal_data_id`) VALUES (1,2),(2,3),(7,8),(8,9),(9,10),(10,11),(11,12),(12,13),(33,44);
/*!40000 ALTER TABLE `organizers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `personal_data`
--

DROP TABLE IF EXISTS `personal_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `personal_data` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `first_name` varchar(40) NOT NULL,
  `last_name` varchar(40) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `personal_data`
--

LOCK TABLES `personal_data` WRITE;
/*!40000 ALTER TABLE `personal_data` DISABLE KEYS */;
INSERT INTO `personal_data` (`id`, `first_name`, `last_name`) VALUES (1,'Pedro','Lopez'),(2,'Lucass','Gonzales'),(3,'Lautaro','Gimenez'),(8,'Camila','Gomez'),(9,'Julieta','De los Santos'),(10,'Fabian','Barrientos'),(11,'Rocio','Manchamelo'),(12,'Agustin','Gonzales'),(13,'Francisco','Sosa'),(26,'Franco','Barrientos'),(42,'franco','barrientos'),(43,'Camila','Gomez'),(44,'franco','barrientos');
/*!40000 ALTER TABLE `personal_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `places`
--

DROP TABLE IF EXISTS `places`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `places` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `address` varchar(40) NOT NULL,
  `address_number` bigint NOT NULL,
  `city` varchar(40) NOT NULL,
  `country` varchar(40) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `places`
--

LOCK TABLES `places` WRITE;
/*!40000 ALTER TABLE `places` DISABLE KEYS */;
INSERT INTO `places` (`id`, `address`, `address_number`, `city`, `country`) VALUES (1,'Colombia',444,'Corrientes','Argentina'),(3,'Rio Juramentos',112,'Corrientes','Argentina'),(8,'Avenida de Mayo',123,'Buenos Aires','Argentina'),(9,'San Martín',456,'San Rafel','Argentina'),(10,'Belgrano',786,'San Rafel','Argentina'),(11,'Libertad',221,'Carloz Paz','Argentina'),(12,'Sarmiento',202,'Carloz Paz','Argentina'),(13,'Mitre',303,'Carloz Paz','Argentina'),(34,'jamaica 4207',4207,'CORRIENTES','Argentina');
/*!40000 ALTER TABLE `places` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_events`
--

DROP TABLE IF EXISTS `user_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_events` (
  `user_id` bigint NOT NULL,
  `event_id` bigint NOT NULL,
  PRIMARY KEY (`user_id`,`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_events`
--

LOCK TABLES `user_events` WRITE;
/*!40000 ALTER TABLE `user_events` DISABLE KEYS */;
INSERT INTO `user_events` (`user_id`, `event_id`) VALUES (1,1),(1,2),(1,8),(1,10),(1,11),(1,12),(1,26),(11,2),(11,7),(11,10);
/*!40000 ALTER TABLE `user_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `personal_data_id` bigint DEFAULT NULL,
  `role` enum('ADMIN','USER') DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `email_2` (`email`),
  KEY `fk_users_personal_data` (`personal_data_id`),
  CONSTRAINT `fk_users_personal_data` FOREIGN KEY (`personal_data_id`) REFERENCES `personal_data` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `email`, `password`, `personal_data_id`, `role`) VALUES (1,'correo@correo.com','$2a$10$Ch0PDhF7y5JibOm4eKbGSunXllhg4HgQMeTzM4gRUMdi1zG22Z/06',1,'USER'),(2,'admin@admin.com','$2a$10$Ch0PDhF7y5JibOm4eKbGSunXllhg4HgQMeTzM4gRUMdi1zG22Z/06',26,'ADMIN'),(10,'francobarrientos@gmail.com','$2a$10$DZItmZoryQLa62pJxZ7yAuTrGShbasNvAGRuaXttqZg7k7HccHxkC',42,'USER');
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

-- Dump completed on 2023-12-11  9:18:35
