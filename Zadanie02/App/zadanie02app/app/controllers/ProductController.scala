package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import play.api.libs.json._
import models.Product

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

    private var productList =  ListBuffer(
        Product(1, "Młotek", 55.0),
        Product(2, "Łopata", 125.0),
        Product(3, "Kilof", 88.0)
    )

// show all
    def getAll() = Action {
        if (productList.isEmpty) {
            NoContent
        }
        else {
            Ok(Json.toJson(productList))
        }
    }

// show by id
    def getById(id: Long) = Action {
        productList.find(p => { p.id == id }) match {
            case Some(product) => Ok(Json.toJson(product))
            case None => NotFound("Brak takiego produktu")
        }
    }

// update
    def update(id: Long) = Action(parse.json) {
        request => request.body.validate[Product].map {
            updatedProduct => val idx = productList.indexWhere(_.id == id)

            if (idx >= 0) {
                productList.update(idx, updatedProduct)
                Ok(Json.toJson(updatedProduct))
            } else {
                NotFound("Brak takiego produktu")
            }
        }.getOrElse {
            BadRequest("Niepoprawny JSON")
        }
    }

// delete
    def delete(id: Long) = Action {
        productList.find(p => { p.id == id }) match {
            case Some(product) => 
                productList -= product
                Ok(Json.toJson(product))
            case None => NotFound("Brak takiego produktu")
        }
    }

// add
    def add() = Action(parse.json) {
        request => request.body.validate[Product].map {
            newProduct => {
                productList += newProduct
                Created(Json.toJson(newProduct))
            }
        }.getOrElse {
            BadRequest("Niepoprawny JSON")
        }
    }

}