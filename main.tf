terraform {
  required_providers {
    cocktails = {
      source  = "hashicorp/cocktails"
      version = "0.1.0"
    }
  }
}

provider "cocktails" {}

resource "cocktails_ingredient" "fruchtsirup_preiselbeere" {
  name = "Fruchtsirup Preiselbeere"
}

resource "cocktails_ingredient" "buttermilch" {
  name = "Buttermilch"
}

resource "cocktails_recipe" "pink_power" {
  name = "Pink Power"

  instruction {
    amount_cl     = 4
    ingredient_id = cocktails_ingredient.fruchtsirup_preiselbeere.id
  }

  instruction {
    amount_cl     = 15
    ingredient_id = cocktails_ingredient.buttermilch.id
  }
}
