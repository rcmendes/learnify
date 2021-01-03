package com.ciriguela.learnify.api.game

import com.ciriguela.learnify.api.game.persistence.Category
import java.time.LocalDateTime
import java.util.*

data class CategoryDTO(
        val uuid: UUID,
        val title: String,
        val description: String? = null,
        val createdAt: LocalDateTime? = null,
        val updatedAt: LocalDateTime? = null) {


    companion object {
        fun from(category: Category): CategoryDTO {
            return CategoryDTO(
                    title = category.title,
                    description = category.description,
                    uuid = category.uuid,
                    createdAt = category.createdAt,
                    updatedAt = category.updatedAt)
        }
    }
}

data class NewCategory(val title: String, val description: String? = null) {
    fun toCategory(): Category {
        return Category(title = this.title, description = this.description, uuid = UUID.randomUUID())
    }
}

//data class Asset(val name: String, val description: String? = null, val url: String, val type: com.ciriguela.learnify.api.game.persistence.AssetType, val category: Category, val databaseInfo: DatabaseInfo)
//
//data class Quiz(val title: String, val description: String? = null, val question: Asset, val answer: Asset, val databaseInfo: DatabaseInfo)
//
//data class Player(val nickname: String, val birthDate: LocalDate, val databaseInfo: DatabaseInfo)
//
//data class LevelQuiz(val quiz: Quiz, val points: Int, val databaseInfo: DatabaseInfo)
//
//data class Stage(val name: String, val description: String? = null, val quizzes: Set<LevelQuiz> = HashSet<LevelQuiz>(), val databaseInfo: DatabaseInfo)
//
//data class Game(val player: Player, val stages: Set<Stage> = HashSet<Stage>(), val databaseInfo: DatabaseInfo)
//
//data class Score(val game: Game, val success: Set<LevelQuiz>, val failures: Set<LevelQuiz>, val databaseInfo: DatabaseInfo)
//
//data class NewAsset(val name: String, val description: String? = null, val url: String, val type: com.ciriguela.learnify.api.game.persistence.AssetType, val category: Category)
