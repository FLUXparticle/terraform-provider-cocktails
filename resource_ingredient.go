package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-cocktails/model"
)

func resourceIngredient() *schema.Resource {
	return &schema.Resource{
		Create: resourceIngredientCreate,
		Read:   resourceIngredientRead,
		Update: resourceIngredientUpdate,
		Delete: resourceIngredientDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIngredientCreate(d *schema.ResourceData, m any) error {
	ingredient := toIngredient(d)

	client := resty.New()
	_, err := client.R().
		SetBody(&ingredient).
		SetResult(&ingredient).
		Post("http://localhost:8080/ingredients")

	if err != nil {
		return fmt.Errorf("error creating ingredient: %w", err)
	}

	return fromIngredient(d, &ingredient)
}

func resourceIngredientRead(d *schema.ResourceData, m any) error {
	return nil
}

func resourceIngredientUpdate(d *schema.ResourceData, m any) error {
	return nil
}

func resourceIngredientDelete(d *schema.ResourceData, m any) error {
	return nil
}

func fromIngredient(d *schema.ResourceData, ingredient *model.Ingredient) error {
	d.SetId(strconv.Itoa(ingredient.IngredientID))

	if err := d.Set("name", ingredient.Name); err != nil {
		return err
	}

	return nil
}

func toIngredient(d *schema.ResourceData) model.Ingredient {
	return model.Ingredient{
		Name: d.Get("name").(string),
	}
}
