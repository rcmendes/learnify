package com.ciriguela.learnify.api.game.persistence

import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.Id
import org.springframework.data.annotation.LastModifiedDate
import org.springframework.data.relational.core.mapping.Table
import java.time.LocalDateTime
import java.util.*
import javax.annotation.processing.Generated

enum class AssetType(val value: String) {
    IMAGE("image"),
    SOUND("sound"),
    TEXT("text")
}

@Table("categories")
data class Category(
        @Id
        val id: Long?=null,
        val uuid: UUID,
        val title: String,
        val description: String? = null,
        @CreatedDate
        val createdAt: LocalDateTime? = null,
        @LastModifiedDate
        val updatedAt: LocalDateTime? = null) {
}