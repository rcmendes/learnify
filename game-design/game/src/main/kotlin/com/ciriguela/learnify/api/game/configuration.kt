package com.ciriguela.learnify.api.game

import io.r2dbc.postgresql.PostgresqlConnectionConfiguration
import io.r2dbc.postgresql.PostgresqlConnectionFactory
import io.r2dbc.spi.ConnectionFactory
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.data.r2dbc.config.AbstractR2dbcConfiguration
import org.springframework.data.r2dbc.core.DatabaseClient
import org.springframework.data.r2dbc.repository.config.EnableR2dbcRepositories

@Configuration
@EnableR2dbcRepositories
class PostgresConfig : AbstractR2dbcConfiguration() {

    @Bean
    override fun connectionFactory(): ConnectionFactory {
        // postgres
        return PostgresqlConnectionFactory(
                PostgresqlConnectionConfiguration.builder()
                        .host("localhost")
                        .database("learnify_game_dev")
                        .username("dev")
                        .password("dev")
                        .build()
        )
    }
//
//    @Bean
//    fun databaseClient(connectionFactory: ConnectionFactory?): DatabaseClient? {
//        return DatabaseClient.builder()
//                .connectionFactory(connectionFactory!!) //.bindMarkers(() -> BindMarkersFactory.named(":", "", 20).create())
//                .namedParameters(true)
//                .build()
//    }
}

