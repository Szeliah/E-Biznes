package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import play.api.libs.json._
import models.Category


@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  
    private val categoryList = ListBuffer(
        Category(1, "Narzędzia"),
        Category(2, "Ubrania robocze"),
        Category(3, "Kaski")
    )

//  show all
    def getAll: Action[AnyContent] = Action {
            if (categoryList.isEmpty) {
            NoContent
        }
        else {
            Ok(Json.toJson(categoryList))
        }
    }

// show by id
    def getById(id: Long) = Action {
        categoryList.find(c => { c.id == id }) match {
            case Some(category) => Ok(Json.toJson(category))
            case None => NotFound("Brak takiej kategorii")
        }
    }

// update
    def update(id: Long) = Action(parse.json) { request =>
        request.body.validate[Category].map {
            updatedCategory => {
                val idx = categoryList.indexWhere(_.id == id)
                if (idx >= 0) {
                    categoryList.update(idx, updatedCategory)
                    Ok(Json.toJson(updatedCategory))
                }
                else {
                    NotFound("Brak takiej kategorii")
                }
            }
        }.getOrElse(BadRequest("Niepoprawny JSON"))
    }
 
//  delete
    def delete(id: Long) = Action {
        categoryList.find(c => { c.id == id }) match {
            case Some(category) => 
                categoryList -= category
                Ok(Json.toJson(category))
            case None => NotFound("Brak takiej kategorii")
        }
    }

// add
    def add() = Action(parse.json) {
       request => request.body.validate[Category].map {
            newCategory => {
                categoryList += newCategory
                Created(Json.toJson(newCategory))
            }
        }.getOrElse {
            BadRequest("Niepoprawny JSON")
        }
    }

}