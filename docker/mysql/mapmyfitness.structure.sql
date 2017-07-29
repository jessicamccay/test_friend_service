DROP TABLE IF EXISTS `friend`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `friend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `from_user_id` bigint(20) NOT NULL,
  `to_user_id` bigint(20) NOT NULL,
  `friends_since` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `friend_pair` (`from_user_id`,`to_user_id`),
  KEY `mmf_friend_to_user_id` (`to_user_id`),
  KEY `mmf_friend_from_user_id` (`from_user_id`),
  CONSTRAINT `from_user_id_refs_id_70e933fe1a6eb01f` FOREIGN KEY (`from_user_id`) REFERENCES `auth_user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `to_user_id_refs_id_70e933fe1a6eb01f` FOREIGN KEY (`to_user_id`) REFERENCES `auth_user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;