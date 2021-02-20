package com.ciriguela.learnify.gamedesign.api.persistence

import org.springframework.data.r2dbc.repository.Query
import org.springframework.data.repository.reactive.ReactiveCrudRepository
import org.springframework.stereotype.Repository
import java.util.*

//interface CategoryRepository {
//    suspend fun insert(category: Category): Category
//    suspend fun findOne(uuid: UUID): Category?
//    suspend fun findAll(): Flow<Category>
//    suspend fun delete(uuid: UUID)
//    suspend fun filter(predicate: (Category) -> Boolean): List<Category>
//}


@Repository
interface CategoryRepository: ReactiveCrudRepository<Category, Long> {
    @Query("select c.* from categories c where c.uuid = :uuid ")
    suspend fun findByUUID(uuid: UUID): Category?
    suspend fun findByTitle(title: String): Category?
}

//@Repository
//class CategoryRepositoryPostgres(private val client: DatabaseClient) : CategoryRepository {
//    override suspend fun insert(category: Category): Category {
//        val result = client.insert().into(Category::class.java).using(category).fetch().awaitOne()
//        return category.copy(createdAt = result["created_at"] as LocalDateTime)
//    }
//
//    override suspend fun findOne(uuid: UUID): Category? {
//        TODO("Not yet implemented")
//    }
//
//    @FlowPreview
//    override suspend fun findAll() = this.client
//            .select()
//            .from(Category::class.java)
//            .`as`(Category::class.java)
//            .all()
//            .asFlow()
//
//
//    override suspend fun delete(uuid: UUID) {
//        TODO("Not yet implemented")
//    }
//
//    override suspend fun filter(predicate: (Category) -> Boolean): List<Category> {
//        TODO("Not yet implemented")
//    }
//
//}

//class CategoryMapper : BiFunction<Row, Any, Category> {
//    override fun apply(row: Row, u: Any): Category {
//        val createdAt = row.get("created_at", LocalDateTime::class.java)!!
//        val updatedAt = row.get("updated_at", LocalDateTime::class.java)
//        val uuid = row.get("uuid", UUID::class.java)!!
//        val title = row.get("title", String::class.java)!!
//        val description = row.get("description", String::class.java)
//
//        val dbInfo = DatabaseInfo(uuid = uuid, createdAt = createdAt, updatedAt = updatedAt)
//        return Category(title = title, description = description, databaseInfo = dbInfo)
//    }
//
//}

//@Component
//class CategoryRepositoryInMemory : CategoryRepository {
//    private val data = HashSet<Category>()
//
//    override suspend fun insert(category: Category): Category {
//        val uuid = UUID.randomUUID()
//        val createdAt = LocalDateTime.now()
//        val storedCategory = category.copy(uuid=uuid, createdAt = createdAt)
//
//        this.data.add(storedCategory)
//
//        return storedCategory
//    }
//
//    override suspend fun findOne(uuid: UUID) = this.filter { it.uuid == uuid }.firstOrNull()?.copy()
//
//    @FlowPreview
//    override suspend fun findAll() = this.data.map { it.copy() }.toList().asFlow()
//
//    override suspend fun delete(uuid: UUID) {
//        this.data.removeIf { it.uuid == uuid }
//    }
//
//    override suspend fun filter(predicate: (Category) -> Boolean): List<Category> {
//        return this.data.filter(predicate)
//    }
//
//    fun clear() {
//        this.data.clear()
//    }
//}


