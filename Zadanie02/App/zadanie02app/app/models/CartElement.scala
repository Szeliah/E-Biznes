package models
import play.api.libs.json._

case class CartElement(id: Long, productId: Long, quantity: Int)

object CartElement {
  implicit val format: OFormat[CartElement] = Json.format[CartElement]
}