---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_theme Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::QuickSight::Theme
---

# awscc_quicksight_theme (Data Source)

Data Source schema for AWS::QuickSight::Theme



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `aws_account_id` (String)
- `base_theme_id` (String)
- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `created_time` (String)
- `last_updated_time` (String)
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `theme_id` (String)
- `type` (String)
- `version` (Attributes) (see [below for nested schema](#nestedatt--version))
- `version_description` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `data_color_palette` (Attributes) (see [below for nested schema](#nestedatt--configuration--data_color_palette))
- `sheet` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet))
- `typography` (Attributes) (see [below for nested schema](#nestedatt--configuration--typography))
- `ui_color_palette` (Attributes) (see [below for nested schema](#nestedatt--configuration--ui_color_palette))

<a id="nestedatt--configuration--data_color_palette"></a>
### Nested Schema for `configuration.data_color_palette`

Read-Only:

- `colors` (List of String)
- `empty_fill_color` (String)
- `min_max_gradient` (List of String)


<a id="nestedatt--configuration--sheet"></a>
### Nested Schema for `configuration.sheet`

Read-Only:

- `tile` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet--tile))
- `tile_layout` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout))

<a id="nestedatt--configuration--sheet--tile"></a>
### Nested Schema for `configuration.sheet.tile`

Read-Only:

- `border` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet--tile--border))

<a id="nestedatt--configuration--sheet--tile--border"></a>
### Nested Schema for `configuration.sheet.tile.border`

Read-Only:

- `show` (Boolean)



<a id="nestedatt--configuration--sheet--tile_layout"></a>
### Nested Schema for `configuration.sheet.tile_layout`

Read-Only:

- `gutter` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout--gutter))
- `margin` (Attributes) (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout--margin))

<a id="nestedatt--configuration--sheet--tile_layout--gutter"></a>
### Nested Schema for `configuration.sheet.tile_layout.margin`

Read-Only:

- `show` (Boolean)


<a id="nestedatt--configuration--sheet--tile_layout--margin"></a>
### Nested Schema for `configuration.sheet.tile_layout.margin`

Read-Only:

- `show` (Boolean)




<a id="nestedatt--configuration--typography"></a>
### Nested Schema for `configuration.typography`

Read-Only:

- `font_families` (Attributes List) (see [below for nested schema](#nestedatt--configuration--typography--font_families))

<a id="nestedatt--configuration--typography--font_families"></a>
### Nested Schema for `configuration.typography.font_families`

Read-Only:

- `font_family` (String)



<a id="nestedatt--configuration--ui_color_palette"></a>
### Nested Schema for `configuration.ui_color_palette`

Read-Only:

- `accent` (String)
- `accent_foreground` (String)
- `danger` (String)
- `danger_foreground` (String)
- `dimension` (String)
- `dimension_foreground` (String)
- `measure` (String)
- `measure_foreground` (String)
- `primary_background` (String)
- `primary_foreground` (String)
- `secondary_background` (String)
- `secondary_foreground` (String)
- `success` (String)
- `success_foreground` (String)
- `warning` (String)
- `warning_foreground` (String)



<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `actions` (List of String)
- `principal` (String)
- `resource` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


<a id="nestedatt--version"></a>
### Nested Schema for `version`

Read-Only:

- `arn` (String)
- `base_theme_id` (String)
- `configuration` (Attributes) (see [below for nested schema](#nestedatt--version--configuration))
- `created_time` (String)
- `description` (String)
- `errors` (Attributes List) (see [below for nested schema](#nestedatt--version--errors))
- `status` (String)
- `version_number` (Number)

<a id="nestedatt--version--configuration"></a>
### Nested Schema for `version.configuration`

Read-Only:

- `data_color_palette` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--data_color_palette))
- `sheet` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet))
- `typography` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--typography))
- `ui_color_palette` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--ui_color_palette))

<a id="nestedatt--version--configuration--data_color_palette"></a>
### Nested Schema for `version.configuration.data_color_palette`

Read-Only:

- `colors` (List of String)
- `empty_fill_color` (String)
- `min_max_gradient` (List of String)


<a id="nestedatt--version--configuration--sheet"></a>
### Nested Schema for `version.configuration.sheet`

Read-Only:

- `tile` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet--tile))
- `tile_layout` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout))

<a id="nestedatt--version--configuration--sheet--tile"></a>
### Nested Schema for `version.configuration.sheet.tile_layout`

Read-Only:

- `border` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout--border))

<a id="nestedatt--version--configuration--sheet--tile_layout--border"></a>
### Nested Schema for `version.configuration.sheet.tile_layout.border`

Read-Only:

- `show` (Boolean)



<a id="nestedatt--version--configuration--sheet--tile_layout"></a>
### Nested Schema for `version.configuration.sheet.tile_layout`

Read-Only:

- `gutter` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout--gutter))
- `margin` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout--margin))

<a id="nestedatt--version--configuration--sheet--tile_layout--gutter"></a>
### Nested Schema for `version.configuration.sheet.tile_layout.gutter`

Read-Only:

- `show` (Boolean)


<a id="nestedatt--version--configuration--sheet--tile_layout--margin"></a>
### Nested Schema for `version.configuration.sheet.tile_layout.margin`

Read-Only:

- `show` (Boolean)




<a id="nestedatt--version--configuration--typography"></a>
### Nested Schema for `version.configuration.typography`

Read-Only:

- `font_families` (Attributes List) (see [below for nested schema](#nestedatt--version--configuration--typography--font_families))

<a id="nestedatt--version--configuration--typography--font_families"></a>
### Nested Schema for `version.configuration.typography.font_families`

Read-Only:

- `font_family` (String)



<a id="nestedatt--version--configuration--ui_color_palette"></a>
### Nested Schema for `version.configuration.ui_color_palette`

Read-Only:

- `accent` (String)
- `accent_foreground` (String)
- `danger` (String)
- `danger_foreground` (String)
- `dimension` (String)
- `dimension_foreground` (String)
- `measure` (String)
- `measure_foreground` (String)
- `primary_background` (String)
- `primary_foreground` (String)
- `secondary_background` (String)
- `secondary_foreground` (String)
- `success` (String)
- `success_foreground` (String)
- `warning` (String)
- `warning_foreground` (String)



<a id="nestedatt--version--errors"></a>
### Nested Schema for `version.errors`

Read-Only:

- `message` (String)
- `type` (String)
