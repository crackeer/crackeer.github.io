<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'http://i.sems-model.api.realsee.com/style/v1/alg_config.json?style_id=1122',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'GET',
  CURLOPT_HTTPHEADER => array(
    'Cookie: lianjia_ssid=85845cbf-aaba-481b-9350-b7f9c9b9a848; lianjia_uuid=b3573a2d-69a8-4409-974b-24fc106b98a2'
  ),
));

$response = curl_exec($curl);

curl_close($curl);

$data = json_decode($response);

var_dump(json_encode($newData["data"]["config"]));

