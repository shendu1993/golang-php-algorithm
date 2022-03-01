<?php


//加载数组case
require_once "ArrayCase.php";
$arrayObj = new ArrayCase();

//https://leetcode.com/problems/two-sum/
//test twoSum
//dd($arrayObj->twoSum([2, 7, 11, 15], 9));
//test twoSumV2
//dd($arrayObj->twoSumV2([2, 7, 11, 15], 9));




/**
 * 打印函数
 * @param $data
 * Author: guanlijian
 * Date: 2022/3/1 12:50
 */
function dd($data)
{
    echo "------------start------------\n";
    var_dump($data);
    echo "------------end--------------\n\n";
    echo "------------分割线------------\n\n";
}

