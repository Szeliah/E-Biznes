package zadanie03

import io.github.cdimascio.dotenv.dotenv
import kotlinx.coroutines.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.websocket.WebSockets
import io.ktor.client.plugins.websocket.webSocket
import io.ktor.websocket.Frame
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.websocket.*
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import kotlinx.serialization.Serializable
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.client.plugins.websocket.*
import kotlinx.serialization.json.buildJsonObject
import kotlinx.serialization.json.put
import kotlinx.serialization.json.putJsonObject
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonNull

@Serializable
data class Message(val content: String)


fun main() {
    
    val dotenv = dotenv()
    val botToken = dotenv["DISCORD_TOKEN"] ?: error("Brak tokena bota w .env")
    val channelId = dotenv["DISCORD_CHANNEL_ID"] ?: error("Brak ID kanału w .env")

 
    val client = HttpClient(CIO) {
        install(io.ktor.client.plugins.websocket.WebSockets)
        install(ContentNegotiation) {json()}
    }



    runBlocking {
        try {
            val res: HttpResponse = client.post("https://discord.com/api/v10/channels/$channelId/messages") {
                header("Authorization", "Bot $botToken")
                contentType(ContentType.Application.Json)
                setBody(Message("Witaj, zadanie03 ebiznes - bot zaczął pracę"))
            }

            if (res.status.isSuccess()) {
                println("Wiadomość wysłana poprawnie")
            } else {
                println("Błąd HTTP: ${res.status}")
            }
        }
        catch (e: Exception) {
            println("Wyjątek: ${e.message}")
        }

        client.close()

    }
}