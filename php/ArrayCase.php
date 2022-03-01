<?php

/**
 * Created by : PhpStorm
 * User: guanlijian
 * Date: 2022/3/1
 * Time: 10:55
 */


class ArrayCase
{
    /**
     * case https://leetcode.com/problems/two-sum/
     * @param array $nums
     * @param int $target
     * @return array
     * Author: guanlijian
     * Date: 2022/3/1 11:01
     */
    public function twoSum(array $nums, int $target): array
    {

        $m = [];
        $k = 0;
        for ($i = 0; $i < count($nums); $i++) {
            //获取当前的值，达到目标需要的差值
            $another = $target - $nums[$i];
            //如果新数组中存在，则直接返回
            $k++;
            if (isset($m[$another])) {
                $this->ddCycleTimes($k);
                return [$m[$another], $i];
            }
            //压入数组中
            $m[$nums[$i]] = $i;
        }
        return [];
    }

    /**
     * case https://leetcode.com/problems/two-sum/
     * @param array $nums
     * @param int $target
     * @return array|int[]
     * Author: guanlijian
     * Date: 2022/3/1 11:08
     */
    public function twoSumV2(array $nums, int $target): array
    {
        $k = 0;
        for ($i = 0; $i < count($nums); $i++) {
            for ($j = $i + 1; $j < count($nums); $j++) {
                $k++;
                if ($nums[$i] + $nums[$j] == $target) {
                    $this->ddCycleTimes($k);
                    return [$i, $j];
                }
            }
        }
        return [];
    }

    public function ddCycleTimes(int $num)
    {
        echo "Cycle " . $num . " Times \n";
    }

}