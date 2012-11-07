package goxsdpkg

//	COLLADA Schema
//	Version 1.4.1 (June 23, 2006)
//	Copyright (C) 2005, 2006 The Khronos Group Inc., Sony Computer Entertainment Inc.
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

type Bool xs:boolean

type DateTime xs:dateTime

type Float xs:double

type Int xs:long

type Name xs:Name

type String xs:string

type Token xs:token

type Uint xs:unsignedLong








type Bool2 ListOfBools

type Bool3 ListOfBools

type Bool4 ListOfBools

type Float2 ListOfFloats

type Float3 ListOfFloats

type Float4 ListOfFloats

type Float7 ListOfFloats

type Float2x2 ListOfFloats

type Float3x3 ListOfFloats

type Float4x4 ListOfFloats

type Float2x3 ListOfFloats

type Float2x4 ListOfFloats

type Float3x2 ListOfFloats

type Float3x4 ListOfFloats

type Float4x2 ListOfFloats

type Float4x3 ListOfFloats

type Int2 ListOfInts

type Int3 ListOfInts

type Int4 ListOfInts

type Int2x2 ListOfInts

type Int3x3 ListOfInts

type Int4x4 ListOfInts

//	An enumuerated type specifying the acceptable morph methods.
type MorphMethodType xs:string

//	An enumerated type specifying the acceptable node types.
type NodeType xs:string

//	This type is used for URI reference which can only reference a resource declared within it's same document.
type URIFragmentType xs:string

//	An enumerated type specifying the acceptable up-axis values.
type UpAxisType xs:string

//	An enumerated type specifying the acceptable document versions.
type VersionType xs:string

type Fx_color_common float4

type Fx_opaque_enum xs:string

type Fx_surface_type_enum xs:string

type Fx_surface_face_enum xs:string

//	The per-texel layout of the format.  The length of the string indicate how many channels there are and the letter respresents the name of the channel.  There are typically 0 to 4 channels.
type Fx_surface_format_hint_channels_enum xs:string

//	Each channel of the texel has a precision.  Typically these are all linked together.  An exact format lay lower the precision of an individual channel but applying a higher precision by linking the channels together may still convey the same information.
type Fx_surface_format_hint_precision_enum xs:string

//	Each channel represents a range of values. Some example ranges are signed or unsigned integers, or between between a clamped range such as 0.0f to 1.0f, or high dynamic range via floating point
type Fx_surface_format_hint_range_enum xs:string

//	Additional hints about data relationships and other things to help the application pick the best format.
type Fx_surface_format_hint_option_enum xs:string

type Fx_sampler_wrap_common xs:NMTOKEN

type Fx_sampler_filter_common xs:NMTOKEN

type Fx_modifier_enum_common xs:NMTOKEN

type Fx_draw_common xs:string

type Fx_pipeline_stage_common xs:string

type GL_MAX_LIGHTS_index xs:nonNegativeInteger

type GL_MAX_CLIP_PLANES_index xs:nonNegativeInteger

type GL_MAX_TEXTURE_IMAGE_UNITS_index xs:nonNegativeInteger

type Gl_blend_type xs:string

type Gl_face_type xs:string

type Gl_blend_equation_type xs:string

type Gl_func_type xs:string

type Gl_stencil_op_type xs:string

type Gl_material_type xs:string

type Gl_fog_type xs:string

type Gl_fog_coord_src_type xs:string

type Gl_front_face_type xs:string

type Gl_light_model_color_control_type xs:string

type Gl_logic_op_type xs:string

type Gl_polygon_mode_type xs:string

type Gl_shade_model_type xs:string

type Gl_alpha_value_type xs:float


type Glsl_float xs:float

type Glsl_int xs:int

type Glsl_bool xs:boolean




type Glsl_bool2 glsl_ListOfBool

type Glsl_bool3 glsl_ListOfBool

type Glsl_bool4 glsl_ListOfBool

type Glsl_float2 glsl_ListOfFloat

type Glsl_float3 glsl_ListOfFloat

type Glsl_float4 glsl_ListOfFloat

type Glsl_float2x2 glsl_ListOfFloat

type Glsl_float3x3 glsl_ListOfFloat

type Glsl_float4x4 glsl_ListOfFloat

type Glsl_int2 glsl_ListOfInt

type Glsl_int3 glsl_ListOfInt

type Glsl_int4 glsl_ListOfInt

type Glsl_pipeline_stage xs:string

type Glsl_identifier xs:token

type Cg_bool xs:boolean

type Cg_float xs:float

type Cg_int xs:int

type Cg_half xs:float

type Cg_fixed xs:float

type Cg_bool1 xs:boolean

type Cg_float1 xs:float

type Cg_int1 xs:int

type Cg_half1 xs:float

type Cg_fixed1 xs:float






type Cg_bool2 cg_ListOfBool

type Cg_bool3 cg_ListOfBool

type Cg_bool4 cg_ListOfBool

type Cg_bool1x1 cg_ListOfBool

type Cg_bool1x2 cg_ListOfBool

type Cg_bool1x3 cg_ListOfBool

type Cg_bool1x4 cg_ListOfBool

type Cg_bool2x1 cg_ListOfBool

type Cg_bool2x2 cg_ListOfBool

type Cg_bool2x3 cg_ListOfBool

type Cg_bool2x4 cg_ListOfBool

type Cg_bool3x1 cg_ListOfBool

type Cg_bool3x2 cg_ListOfBool

type Cg_bool3x3 cg_ListOfBool

type Cg_bool3x4 cg_ListOfBool

type Cg_bool4x1 cg_ListOfBool

type Cg_bool4x2 cg_ListOfBool

type Cg_bool4x3 cg_ListOfBool

type Cg_bool4x4 cg_ListOfBool

type Cg_float2 cg_ListOfFloat

type Cg_float3 cg_ListOfFloat

type Cg_float4 cg_ListOfFloat

type Cg_float1x1 cg_ListOfFloat

type Cg_float1x2 cg_ListOfFloat

type Cg_float1x3 cg_ListOfFloat

type Cg_float1x4 cg_ListOfFloat

type Cg_float2x1 cg_ListOfFloat

type Cg_float2x2 cg_ListOfFloat

type Cg_float2x3 cg_ListOfFloat

type Cg_float2x4 cg_ListOfFloat

type Cg_float3x1 cg_ListOfFloat

type Cg_float3x2 cg_ListOfFloat

type Cg_float3x3 cg_ListOfFloat

type Cg_float3x4 cg_ListOfFloat

type Cg_float4x1 cg_ListOfFloat

type Cg_float4x2 cg_ListOfFloat

type Cg_float4x3 cg_ListOfFloat

type Cg_float4x4 cg_ListOfFloat

type Cg_int2 cg_ListOfInt

type Cg_int3 cg_ListOfInt

type Cg_int4 cg_ListOfInt

type Cg_int1x1 cg_ListOfInt

type Cg_int1x2 cg_ListOfInt

type Cg_int1x3 cg_ListOfInt

type Cg_int1x4 cg_ListOfInt

type Cg_int2x1 cg_ListOfInt

type Cg_int2x2 cg_ListOfInt

type Cg_int2x3 cg_ListOfInt

type Cg_int2x4 cg_ListOfInt

type Cg_int3x1 cg_ListOfInt

type Cg_int3x2 cg_ListOfInt

type Cg_int3x3 cg_ListOfInt

type Cg_int3x4 cg_ListOfInt

type Cg_int4x1 cg_ListOfInt

type Cg_int4x2 cg_ListOfInt

type Cg_int4x3 cg_ListOfInt

type Cg_int4x4 cg_ListOfInt

type Cg_half2 cg_ListOfHalf

type Cg_half3 cg_ListOfHalf

type Cg_half4 cg_ListOfHalf

type Cg_half1x1 cg_ListOfHalf

type Cg_half1x2 cg_ListOfHalf

type Cg_half1x3 cg_ListOfHalf

type Cg_half1x4 cg_ListOfHalf

type Cg_half2x1 cg_ListOfHalf

type Cg_half2x2 cg_ListOfHalf

type Cg_half2x3 cg_ListOfHalf

type Cg_half2x4 cg_ListOfHalf

type Cg_half3x1 cg_ListOfHalf

type Cg_half3x2 cg_ListOfHalf

type Cg_half3x3 cg_ListOfHalf

type Cg_half3x4 cg_ListOfHalf

type Cg_half4x1 cg_ListOfHalf

type Cg_half4x2 cg_ListOfHalf

type Cg_half4x3 cg_ListOfHalf

type Cg_half4x4 cg_ListOfHalf

type Cg_fixed2 cg_ListOfFixed

type Cg_fixed3 cg_ListOfFixed

type Cg_fixed4 cg_ListOfFixed

type Cg_fixed1x1 cg_ListOfFixed

type Cg_fixed1x2 cg_ListOfFixed

type Cg_fixed1x3 cg_ListOfFixed

type Cg_fixed1x4 cg_ListOfFixed

type Cg_fixed2x1 cg_ListOfFixed

type Cg_fixed2x2 cg_ListOfFixed

type Cg_fixed2x3 cg_ListOfFixed

type Cg_fixed2x4 cg_ListOfFixed

type Cg_fixed3x1 cg_ListOfFixed

type Cg_fixed3x2 cg_ListOfFixed

type Cg_fixed3x3 cg_ListOfFixed

type Cg_fixed3x4 cg_ListOfFixed

type Cg_fixed4x1 cg_ListOfFixed

type Cg_fixed4x2 cg_ListOfFixed

type Cg_fixed4x3 cg_ListOfFixed

type Cg_fixed4x4 cg_ListOfFixed

type Cg_pipeline_stage xs:string

type Cg_identifier xs:token

type GLES_MAX_LIGHTS_index xs:nonNegativeInteger

type GLES_MAX_CLIP_PLANES_index xs:nonNegativeInteger

type GLES_MAX_TEXTURE_COORDS_index xs:nonNegativeInteger

type GLES_MAX_TEXTURE_IMAGE_UNITS_index xs:nonNegativeInteger

type Gles_texenv_mode_enums xs:token

type Gles_texcombiner_operatorRGB_enums xs:token

type Gles_texcombiner_operatorAlpha_enums xs:token

type Gles_texcombiner_source_enums xs:token

type Gles_texcombiner_operandRGB_enums gl_blend_type

type Gles_texcombiner_operandAlpha_enums gl_blend_type

type Gles_texcombiner_argument_index_type xs:nonNegativeInteger

type Gles_sampler_wrap xs:NMTOKEN

type Gles_stencil_op_type xs:string


type Gles_rendertarget_common xs:NCName

type SpringType xs:NMTOKEN

type Common_profile_input xs:NMTOKEN

type Common_profile_param xs:NMTOKEN
