package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import play.api.libs.json._
import models.CartElement

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  
private var cartElements = ListBuffer[CartElement]()

//  show all
    def getAll: Action[AnyContent] = Action {
            if (cartElements.isEmpty) {
            NoContent
        }
        else {
            Ok(Json.toJson(cartElements))
        }
    }

// show by id
    def getById(id: Long) = Action {
        cartElements.find(c => { c.id == id }) match {
            case Some(element) => Ok(Json.toJson(element))
            case None => NotFound("Brak takiego przedmoiotu w koszyku")
        }
    }

// update
    def update(id: Long) = Action(parse.json) { request =>
        request.body.validate[CartElement].map {
            updatedCart => {
                val idx = cartElements.indexWhere(_.id == id)
                if (idx >= 0) {
                    cartElements.update(idx, updatedCart)
                    Ok(Json.toJson(updatedCart))
                }
                else {
                    NotFound("Brak takiego przedmoiotu w koszyku")
                }
            }
        }.getOrElse(BadRequest("Niepoprawny JSON"))
    }
 
//  delete
    def delete(id: Long) = Action {
        cartElements.find(c => { c.id == id }) match {
            case Some(cart) => 
                cartElements -= cart
                Ok(Json.toJson(cart))
            case None => NotFound("Brak takiej kategorii")
        }
    }

// add
    def add() = Action(parse.json) {
       request => request.body.validate[CartElement].map {
            newCart => {
                cartElements += newCart
                Created(Json.toJson(newCart))
            }
        }.getOrElse {
            BadRequest("Niepoprawny JSON")
        }
    }

}