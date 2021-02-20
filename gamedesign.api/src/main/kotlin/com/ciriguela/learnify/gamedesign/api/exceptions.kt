package com.ciriguela.learnify.gamedesign.api


data class APIError(val code: String, val message: String)

sealed class DomainError(val code: String, message: String) : Exception(message)
sealed class EntityAlreadyExistsError(code: String, message: String) : DomainError(code, message)

class CategoryAlreadyExistsError(private val propertyNameAndValue: Pair<String, String>) :
        EntityAlreadyExistsError(
                code = "category-already-exists",
                message = "A Category with ${propertyNameAndValue.first} " +
                           "'${propertyNameAndValue.second}' already exists.")


