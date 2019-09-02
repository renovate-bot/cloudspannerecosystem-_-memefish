select 1 + 2, 1 - 2,
       1 * 2, 2 / 2,
       +1++1, -1+-1,
       1 ^ 2, 2 & 1, 2 | 1,
       1 << 2, 2 >> 1,
       foo.bar,
       (select 1 `1`).1,
       [1, 2, 3][offset(1)],
       [1, 2, 3][`offset`(1)],
       [1, 2, 3][ordinal(1)],
       case
       when 1 = 1 then "1 = 1"
       else            "else"
       end,
       case 1
       when 1 then "1"
       when 2 then "2"
       else        "other"
       end,
       date_add(date "2019-09-01", interval 5 day),
       timestamp_add(timestamp "2019-09-01 08:11:22", interval 5 hour)
