package com.ciriguela.learnify.api.game

import com.ciriguela.learnify.api.game.persistence.CategoryRepositoryInMemory
import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test

internal class CategoryServiceTest {
    companion object {
        val categoryRepo = CategoryRepositoryInMemory()
        val service = CategoryService(categoryRepo)
    }

    @Test
    @DisplayName("Creation of a category and persisting into repository")
    fun createCategory() = runBlocking {
        val title = "Title"
        val description = "Description"

        val newCategory = NewCategory(title = title, description = description)


        val category = service.createCategory(newCategory)
        assert(category.title == title)
        assert(category.description == description)
        assert(category.databaseInfo.id > 0)
        assert(category.databaseInfo.updatedAt == null)
    }

    @Test
    fun listCategories() = runBlocking {
        val list = (1..10)
                .map { NewCategory(title = "Category $it", description = "Description $it") }
                .map {
                    service.createCategory(it)
                }

        assert(list.isNotEmpty())
        assert(list.size == 10)

    }

    @Test
    fun findCategory() = runBlocking {
        val list = (1..10)
                .map { NewCategory(title = "Category $it", description = "Description $it") }
                .map {
                    service.createCategory(it)
                }

        list.forEach {
            val found = service.findCategory(it.databaseInfo.id)
            assert(found != null)
            assert(it == found)
        }

        list.forEach {
            val found = service.findCategory(it.databaseInfo.uuid)
            assert(found != null)
            assert(it == found)
        }
    }

    @BeforeEach
    fun setUp() {
        categoryRepo.clear()
    }
}