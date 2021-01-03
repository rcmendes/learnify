package com.ciriguela.learnify.gamedesign.api

import com.ciriguela.learnify.gamedesign.api.persistence.Category
import com.ciriguela.learnify.gamedesign.api.persistence.CategoryRepository
import kotlinx.coroutines.FlowPreview
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.reactive.asFlow
import kotlinx.coroutines.reactive.awaitFirst
import org.springframework.stereotype.Service
import java.util.*

@Service
class CategoryService(val categoryRepository: CategoryRepository) {
    suspend fun createCategory(newCategory: NewCategory): Category {
//        this.categoryRepository.findByTitle(newCategory.title)?:
        val category = newCategory.toCategory()
        return this.categoryRepository.save(category).awaitFirst()
    }

    @FlowPreview
    suspend fun listCategories(): Flow<Category> {
        return this.categoryRepository.findAll().asFlow()
    }

    suspend fun findCategory(uuid: UUID): CategoryDTO? {
        return categoryRepository.findByUUID(uuid)?.let {
            return CategoryDTO.from(it)
        }
    }

}