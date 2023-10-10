package main

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Marshal ...
func Marshal(dest interface{}) ([]byte, error) {
	return json.Marshal(dest)
}

// Unmarshal ...
func Unmarshal(input []byte, data interface{}) error {
	return json.Unmarshal(input, &data)
}

// MapRemoveNumber ...
func MapRemoveNumber(workData map[string]interface{}) map[string]interface{} {
	if workData == nil {
		return nil
	}

	retData := map[string]interface{}{}
	for k, v := range workData {
		retData[k] = ToZero(v)
	}
	return retData
}

// ToZero ...
func ToZero(data interface{}) interface{} {

	if data == nil {
		return nil
	}

	if list, ok := data.([]interface{}); ok {
		retData := []interface{}{}
		for _, item := range list {
			retData = append(retData, ToZero(item))
		}
		return retData
	}

	if list, ok := data.([]map[string]interface{}); ok {
		retData := []interface{}{}
		for _, item := range list {
			retData = append(retData, ToZero(item))
		}
		return retData
	}

	if mapData, ok := data.(map[string]interface{}); ok {
		retData := map[string]interface{}{}
		for k, v := range mapData {
			retData[k] = ToZero(v)
		}
		return retData
	}

	dataType := reflect.TypeOf(data)
	zeroDataType := []string{
		reflect.Int8.String(), reflect.Int16.String(), reflect.Int32.String(), reflect.Int64.String(), reflect.Uint.String(), reflect.Int.String(), reflect.Uint8.String(), reflect.Uint16.String(), reflect.Uint32.String(), reflect.Uint64.String(),
		reflect.Float32.String(), reflect.Float64.String(),
	}

	for _, v := range zeroDataType {
		if v == dataType.String() {
			return 0
		}
	}

	return data
}

var jsonData string = `{"base_url":"https://vr.ziroom.com/vr/1657696069961/4072704/","create_time":"2022-06-12T19:18:07+08:00","model":{"file_url":"model/cyclops-GAv7ZjjKG1WZWxpE.at3d","material_base_url":null,"material_textures":["texture_0.jpg","texture_1.jpg","texture_2.jpg"],"modify_time":"2022-06-12T19:34:17+08:00","type":0},"observers":[{"accessible_nodes":[1,2,3,6,7],"floor_index":0,"index":0,"offset_point_count":0,"position":[1.8830499649047852,0.17,-2.095020055770874],"quaternion":{"w":0.008731637499547295,"x":7.498513074866031e-33,"y":-0.9999618821031793,"z":1.2246001181870215e-16},"standing_position":[1.8830499649047852,-1.3311873735178945,-2.095020055770874],"visible_nodes":[1,2,3,6,7]},{"accessible_nodes":[0,2,3,6,9,10,11,12],"floor_index":0,"index":1,"offset_point_count":0,"position":[1.882449984550476,0.17894096028059722,-0.7170739769935608],"quaternion":{"w":-0.025825554029677237,"x":-0.0013614529721375933,"y":-0.9996602777755719,"z":-0.003242277026393596},"standing_position":[1.882449984550476,-1.3307247872961012,-0.7170739769935608],"visible_nodes":[0,2,3,6,9,10,11,12]},{"accessible_nodes":[0,1,3,4,5],"floor_index":0,"index":2,"offset_point_count":0,"position":[2.0387699604034424,0.15930490016937257,2.297529935836792],"quaternion":{"w":-0.03654714223661472,"x":-0.0006115283349247825,"y":-0.9993263076385644,"z":-0.003297785552801522},"standing_position":[2.0387699604034424,-1.3306436703329136,2.297529935836792],"visible_nodes":[0,1,3,4,5]},{"accessible_nodes":[0,1,2],"floor_index":0,"index":3,"offset_point_count":0,"position":[2.1114299297332764,0.13695090085268022,5.578700065612793],"quaternion":{"w":-0.03275090925401649,"x":-0.0016918035368158623,"y":-0.9994539319633519,"z":-0.004045515583630641},"standing_position":[2.1114299297332764,-1.3302500365902619,5.578700065612793],"visible_nodes":[0,1,2]},{"accessible_nodes":[2,5],"floor_index":0,"index":4,"offset_point_count":0,"position":[0.4415690004825592,0.16918725201860071,2.302229881286621],"quaternion":{"w":-0.7947124623318962,"x":-0.0036990055651711716,"y":-0.6069545849632668,"z":-0.004952621888347388},"standing_position":[0.4415690004825592,-1.3299991697046813,2.302229881286621],"visible_nodes":[2,5]},{"accessible_nodes":[2,4],"floor_index":0,"index":5,"offset_point_count":0,"position":[-2.17684006690979,0.16301083013415338,2.454479932785034],"quaternion":{"w":-0.7628388797705961,"x":-0.0034425576413478418,"y":-0.646559059807801,"z":-0.005134227808822177},"standing_position":[-2.17684006690979,-1.3304054413461948,2.454479932785034],"visible_nodes":[2,4]},{"accessible_nodes":[0,1,7,9,10,11,12],"floor_index":0,"index":6,"offset_point_count":0,"position":[-0.11009199917316437,0.19554219961166383,-0.7679719924926758],"quaternion":{"w":-0.7315547928899536,"x":-0.0022161907700936623,"y":-0.6817788076930082,"z":0.0005650840987978786},"standing_position":[-0.11009199917316437,-1.3305576645566923,-0.7679719924926758],"visible_nodes":[0,1,7,9,10,11,12]},{"accessible_nodes":[0,6,8,9,10],"floor_index":0,"index":7,"offset_point_count":0,"position":[-0.08624009788036346,0.21930970072746278,-0.049502599984407425],"quaternion":{"w":-0.1781836195207241,"x":-0.007394231940660549,"y":-0.9839694686432949,"z":-0.0001545090441634976},"standing_position":[-0.08624009788036346,-1.3299973241834098,-0.049502599984407425],"visible_nodes":[0,6,8,9,10]},{"accessible_nodes":[7],"floor_index":0,"index":8,"offset_point_count":0,"position":[-1.6563199758529663,0.22887759849429132,0.5843020081520081],"quaternion":{"w":-0.6044145391745535,"x":0.0071733976072566345,"y":-0.7965425283952116,"z":0.012313372638282652},"standing_position":[-1.6563199758529663,-1.330159915877742,0.5843020081520081],"visible_nodes":[7]},{"accessible_nodes":[1,6,7,10,11],"floor_index":0,"index":9,"offset_point_count":0,"position":[-0.4493809938430786,0.21628520086407663,-1.4461599588394165],"quaternion":{"w":-0.9794245246800367,"x":-0.0008222534476449962,"y":0.2017978401006683,"z":0.0021341814984422226},"standing_position":[-0.4493809938430786,-1.3310777058564416,-1.4461599588394165],"visible_nodes":[1,6,7,10,11]},{"accessible_nodes":[1,6,7,9,11,12],"floor_index":0,"index":10,"offset_point_count":0,"position":[-0.7600039839744568,0.1906199996173382,-1.1758899688720703],"quaternion":{"w":-0.7836594460904323,"x":-0.0007011993655561965,"y":-0.6211900816292951,"z":-0.0005247475835910507},"standing_position":[-0.7600039839744568,-1.3302897531415279,-1.1758899688720703],"visible_nodes":[1,6,7,9,11,12]},{"accessible_nodes":[1,6,9,10,12],"floor_index":0,"index":11,"offset_point_count":0,"position":[-1.6307799816131592,0.1941855997592211,-0.9025710225105286],"quaternion":{"w":-0.7627064439150975,"x":-0.0006834787813303143,"y":-0.6467442407996731,"z":-0.0005476291091873649},"standing_position":[-1.6307799816131592,-1.3308154885084407,-0.9025710225105286],"visible_nodes":[1,6,9,10,12]},{"accessible_nodes":[1,6,10,11],"floor_index":0,"index":12,"offset_point_count":0,"position":[-3.703779935836792,0.20974080085754396,-1.0182199478149414],"quaternion":{"w":-0.7911340605941365,"x":-0.0007075149862302202,"y":-0.6116421474954813,"z":-0.0005162054325231264},"standing_position":[-3.703779935836792,-1.3306066585609135,-1.0182199478149414],"visible_nodes":[1,6,10,11]}],"panorama":{"count":13,"list":[{"back":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_b.jpg","down":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_d.jpg","front":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_f.jpg","index":0,"left":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_l.jpg","right":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_r.jpg","up":"images/cube_2048/0/e6d15fddbe5ef75c1e559735962e558e/0_u.jpg"},{"back":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_b.jpg","down":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_d.jpg","front":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_f.jpg","index":1,"left":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_l.jpg","right":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_r.jpg","up":"images/cube_2048/1/273b5128bfd784d2df000572ea005595/1_u.jpg"},{"back":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_b.jpg","down":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_d.jpg","front":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_f.jpg","index":2,"left":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_l.jpg","right":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_r.jpg","up":"images/cube_2048/2/530945952a020c732d1d9b4d95030376/2_u.jpg"},{"back":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_b.jpg","down":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_d.jpg","front":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_f.jpg","index":3,"left":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_l.jpg","right":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_r.jpg","up":"images/cube_2048/3/6edf1689ea952d0836bab1412c5c3e29/3_u.jpg"},{"back":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_b.jpg","down":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_d.jpg","front":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_f.jpg","index":4,"left":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_l.jpg","right":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_r.jpg","up":"images/cube_2048/4/3437046ac3006f0ceb0c26a0e13c0e08/4_u.jpg"},{"back":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_b.jpg","down":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_d.jpg","front":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_f.jpg","index":5,"left":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_l.jpg","right":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_r.jpg","up":"images/cube_2048/5/663ad97fff309c34e75e679ca1a8911e/5_u.jpg"},{"back":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_b.jpg","down":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_d.jpg","front":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_f.jpg","index":6,"left":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_l.jpg","right":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_r.jpg","up":"images/cube_2048/6/b216da94d613dedc7b81a1e28d7ef446/6_u.jpg"},{"back":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_b.jpg","down":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_d.jpg","front":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_f.jpg","index":7,"left":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_l.jpg","right":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_r.jpg","up":"images/cube_2048/7/d7dac81fb8aa3eb87bb551be9a02795a/7_u.jpg"},{"back":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_b.jpg","down":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_d.jpg","front":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_f.jpg","index":8,"left":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_l.jpg","right":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_r.jpg","up":"images/cube_2048/8/45bdb309f2acabf99ab9cefbf0f6de09/8_u.jpg"},{"back":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_b.jpg","down":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_d.jpg","front":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_f.jpg","index":9,"left":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_l.jpg","right":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_r.jpg","up":"images/cube_2048/9/b7e85a9f8b13454df0ef592d5b5eea2c/9_u.jpg"},{"back":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_b.jpg","down":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_d.jpg","front":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_f.jpg","index":10,"left":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_l.jpg","right":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_r.jpg","up":"images/cube_2048/10/aeb62d8adfdf34584537e957befd6808/10_u.jpg"},{"back":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_b.jpg","down":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_d.jpg","front":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_f.jpg","index":11,"left":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_l.jpg","right":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_r.jpg","up":"images/cube_2048/11/4fd7207e33a15f10106e49dd61dead97/11_u.jpg"},{"back":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_b.jpg","down":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_d.jpg","front":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_f.jpg","index":12,"left":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_l.jpg","right":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_r.jpg","up":"images/cube_2048/12/a5430dad8f865cd6e4b246da20a88123/12_u.jpg"}]},"picture_url":"","title_picture_url":"","vr_code":"80doNP9KXEdWjDwjJr","vr_type":"reality"}`

func main() {
	workData := map[string]interface{}{}
	Unmarshal([]byte(jsonData), &workData)

	data := MapRemoveNumber(workData)
	bytes, _ := Marshal(data)
	fmt.Println(string(bytes))
}
