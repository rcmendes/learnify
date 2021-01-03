package com.ciriguela.learnify.api.game

import kotlinx.coroutines.FlowPreview
import kotlinx.coroutines.reactive.awaitFirst
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import org.springframework.web.util.UriComponentsBuilder
import java.net.URI
import java.util.*


@RestController
@RequestMapping("/categories")
class CategoryController {

    @Autowired
    lateinit var categoryService: CategoryService

    @FlowPreview
    @GetMapping("/")
    suspend fun listCategories() = categoryService.listCategories()

    @GetMapping("/{uuid}")
    suspend fun getCategoryByUUID(@PathVariable uuid: UUID): ResponseEntity<CategoryDTO?> {
        val category = categoryService.findCategory(uuid) ?: return ResponseEntity.notFound().build()

        return ResponseEntity.ok(category)
    }

    @PostMapping("/")
    @ResponseBody
    suspend fun createCategory(@RequestBody category: NewCategory, uriBuilder: UriComponentsBuilder): ResponseEntity<Any> {
        categoryService.createCategory(category).let {
            val location = uriBuilder
                    .path("/categories/${it.uuid}").build().toUri()

            return ResponseEntity.created(location).build()
        }
    }
}

@ControllerAdvice
internal class GlobalControllerErrorHandler {
    companion object {
        val logger = LoggerFactory.getLogger(GlobalControllerErrorHandler::class.qualifiedName)!!
    }

    @ExceptionHandler(EntityAlreadyExistsError::class)
    fun handleBadRequest(exception: EntityAlreadyExistsError): ResponseEntity<APIError> {
        logger.debug("${exception.toString()}")
        APIError(code = exception.code, message = exception.message!!).let {
            return ResponseEntity(it, HttpStatus.BAD_REQUEST)
        }
    }
}