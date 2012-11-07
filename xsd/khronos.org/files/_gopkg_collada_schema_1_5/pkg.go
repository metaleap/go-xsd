package goxsdpkg

//	COLLADA Schema
//	Version 1.5.0 (August 6, 2008)
//	Copyright (C) 2008 The Khronos Group Inc., Sony Computer Entertainment Inc.
//	All Rights Reserved.
//	Khronos is a trademark of The Khronos Group Inc.
//	COLLADA is a trademark of Sony Computer Entertainment Inc. used by permission by Khronos.
//	Note that this software document is distributed on an "AS IS" basis, with ALL EXPRESS AND
//	IMPLIED WARRANTIES AND CONDITIONS DISCLAIMED, INCLUDING, WITHOUT LIMITATION, ANY IMPLIED
//	WARRANTIES AND CONDITIONS OF MERCHANTABILITY, SATISFACTORY QUALITY, FITNESS FOR A PARTICULAR
//	PURPOSE, AND NON-INFRINGEMENT.
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type Float_type xs:double

type Int_type xs:long

type Uint_type xs:unsignedLong

type Sidref_type xs:string

type Sid_type xs:NCName










type Bool2_type list_of_bools_type

type Bool3_type list_of_bools_type

type Bool4_type list_of_bools_type

type Bool2x2_type list_of_bools_type

type Bool2x3_type list_of_bools_type

type Bool2x4_type list_of_bools_type

type Bool3x2_type list_of_bools_type

type Bool3x3_type list_of_bools_type

type Bool3x4_type list_of_bools_type

type Bool4x2_type list_of_bools_type

type Bool4x3_type list_of_bools_type

type Bool4x4_type list_of_bools_type

type Float2_type list_of_floats_type

type Float3_type list_of_floats_type

type Float4_type list_of_floats_type

type Float7_type list_of_floats_type

type Float2x2_type list_of_floats_type

type Float2x3_type list_of_floats_type

type Float2x4_type list_of_floats_type

type Float3x2_type list_of_floats_type

type Float3x3_type list_of_floats_type

type Float3x4_type list_of_floats_type

type Float4x2_type list_of_floats_type

type Float4x3_type list_of_floats_type

type Float4x4_type list_of_floats_type

type Int2_type list_of_ints_type

type Int3_type list_of_ints_type

type Int4_type list_of_ints_type

type Int2x2_type list_of_ints_type

type Int2x3_type list_of_ints_type

type Int2x4_type list_of_ints_type

type Int3x2_type list_of_ints_type

type Int3x3_type list_of_ints_type

type Int3x4_type list_of_ints_type

type Int4x2_type list_of_ints_type

type Int4x3_type list_of_ints_type

type Int4x4_type list_of_ints_type

type Digits_type xs:unsignedByte

type Magnitude_type xs:short

//	An enumuerated type specifying the acceptable morph methods.
type Morph_method_enum xs:string

//	An enumerated type specifying the acceptable node types.
type Node_enum xs:string

//	An enumerated type specifying the acceptable sampler pre and post behavior attribute types.
type Sampler_behavior_enum xs:string

//	This urifragment_type element is used for URI reference which can only reference a resource declared within it's same document.
type Urifragment_type xs:string

//	An enumerated type specifying the acceptable up-axis values.
type Up_axis_enum xs:string

//	An enumerated type specifying the acceptable document versions.
type Version_enum xs:string

type Image_face_enum xs:string

//	The per-texel layout of the format.  The length of the string indicate how many channels there are and the letter respresents the name of the channel.  There are typically 0 to 4 channels.
type Image_format_hint_channels_enum xs:string

//	Each channel of the texel has a precision.  Typically these are all linked together.  An exact format lay lower the precision of an individual channel but applying a higher precision by linking the channels together may still convey the same information.
type Image_format_hint_precision_enum xs:string

//	Each channel represents a range of values. Some example ranges are signed or unsigned integers, or between between a clamped range such as 0.0f to 1.0f, or high dynamic range via floating point
type Image_format_hint_range_enum xs:string

//	The legal values for the mode attribute on the altitute element in a
//	geographic_location element.
type Altitude_mode_enum xs:string

type Fx_color_type float4_type

type Fx_opaque_enum xs:string

type Fx_sampler_wrap_enum xs:NMTOKEN

type Fx_sampler_min_filter_enum xs:NMTOKEN

type Fx_sampler_mag_filter_enum xs:NMTOKEN

type Fx_sampler_mip_filter_enum xs:NMTOKEN

type Fx_modifier_enum xs:NMTOKEN

//	? GEOMETRY: [default] The geometry associated with this instance_geometry or nstance_material.
//	? SCENE_GEOMETRY: Draw the entire scene's geometry but with this effect, not the effects or
//	materials already associated with the geometry. This is for techniques such as shadow-buffer
//	generation, where you might be interested only in extracting the Z value from the light. This is
//	without regard to ordering on the assumption that ZBuffer handles order.
//	? SCENE_IMAGE: Draw the entire scene naturally into the render targets. Hense producing an image
//	of the scene.  When used more then once, the later uses must include what has been rendered so far
//	to the backbuffer since the first render.  This is for effects that need an accurate image of the scene
//	to work on for effects such as postprocessing blurs.
//	? FULL_SCREEN_QUAD: Positions are 0,0 to 1,1 and the UVs match.
type Fx_draw_type xs:string

type Fx_pipeline_stage_enum xs:string

type Gl_max_lights_index_type xs:nonNegativeInteger

type Gl_max_clip_planes_index_type xs:nonNegativeInteger

type Gl_max_texture_image_units_index_type xs:nonNegativeInteger

type Gl_blend_enum xs:string

type Gl_face_enum xs:string

type Gl_blend_equation_enum xs:string

type Gl_func_enum xs:string

type Gl_stencil_op_enum xs:string

type Gl_material_enum xs:string

type Gl_fog_enum xs:string

type Gl_fog_coord_src_enum xs:string

type Gl_front_face_enum xs:string

type Gl_light_model_color_control_enum xs:string

type Gl_logic_op_enum xs:string

type Gl_polygon_mode_enum xs:string

type Gl_shade_model_enum xs:string

type Gl_alpha_value_type xs:float


type Gles_max_lights_index_type xs:nonNegativeInteger

type Gles_max_clip_planes_index_type xs:nonNegativeInteger

type Gles_max_texture_coords_index_type xs:nonNegativeInteger

type Gles_max_texture_image_units_index_type xs:nonNegativeInteger

type Gles_texenv_mode_enum xs:token

type Gles_texcombiner_operator_rgb_enum xs:token

type Gles_texcombiner_operator_alpha_enum xs:token

type Gles_texcombiner_source_enum xs:token

type Gles_texcombiner_operand_rgb_enum gl_blend_enum

type Gles_texcombiner_operand_alpha_enum gl_blend_enum

type Gles_texcombiner_argument_index_type xs:nonNegativeInteger

type Gles_sampler_wrap_enum xs:NMTOKEN

type Gles_stencil_op_enum xs:string


type Spring_enum xs:NMTOKEN

type Common_profile_input_enum xs:NMTOKEN

type Common_profile_param_enum xs:NMTOKEN

type Dynamic_limit_type float2_type
