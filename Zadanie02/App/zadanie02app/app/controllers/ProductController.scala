package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import models.Product

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

    private val productList =  List(
        Product(1, "Młotek", 55.0),
        Product(2, "Łopata", 125.0),
        Product(3, "Kilof", 88.0)
    )

}