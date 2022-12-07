#!/bin/bash
declare letters=ijk1bl_2fcmNr_v3g_dpnAOC_sQ_JwE4h1a_eMqu_ozB_PID_9Lt_y_H8KxG7F65
declare -i db=62#1lVH34yFmBR
while read items; do
  declare -i len="${#items}-1"
  declare -ia bag=(0 0)
  for i in $(seq 0 $len); do
    declare item="${items:i:1}"
    let "bag[i <= len / 2] |= (1 << (62#$item - 10))"
  done
  declare -i idx='((bag[0] & bag[1]) * db) >> 58'
  let 'idx += 64 * (idx < 0), idx &= 63'
  echo "$((62#${letters:idx:1}))"
done
