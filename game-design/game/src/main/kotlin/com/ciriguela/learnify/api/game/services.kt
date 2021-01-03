package com.ciriguela.learnify.api.game

import com.ciriguela.learnify.api.game.persistence.Category
import com.ciriguela.learnify.api.game.persistence.CategoryRepository
import kotlinx.coroutines.FlowPreview
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.reactive.asFlow
import kotlinx.coroutines.reactive.awaitFirst
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*

@Service
class CategoryService(val categoryRepository: CategoryRepository) {
    suspend fun createCategory(newCategory: NewCategory): Category {
        this.categoryRepository.findByTitle(newCategory.title)?:
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