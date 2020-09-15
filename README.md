# DynamicLevelCacheDispatcher

> **!!!NOTE** this proj is in early deving, be carefull to use it.

## What's this?

Simply, **DynamicLevelCache** is a cache dispatch system, you use it like a cache.  With dynamic calculating , it will determine destination and duration an object stores.

## How it works?

**DynamicLevelCache** itself does not privde data storage. Instead of this, it use exported apis to register cache handler.

## Classic usage
1. data load from database
2. first level cache in mem
3. second level cache in redis


## Features
1. First design for single machine
2. Can work as a (http/rpc)service
3. Use protolbuf as data trans protocol
4. Thread(routine) safe


## Algorithm Desc 
1. with input parameter : getter calling, data fetcher
2. give output result : store level , store duration


## Algorithm Type
1. Destination type:  Avg Split / Weight Split
2. Duration type: Fixed duration / Count weight refering duration / Score weight  refering duration


## Algorithm Implementation
1. We simply split reformat input to time density
2. We score and classify result of first step
3. With score we determine destination and duration for this data

#### score calculatation
1. With data featching duration d1 and duration between two getter calling d2 we get time destiny s=d1/d2. 
2. With s1 = avg(s_in_10_sec), s2=agv(s_in_one_min), s3=avg(s_in_five_min)
3. Score result is r = s1*10+s2*5*s3*3,  (weight 10/5/3 is global parameter)
4. We rank score for all data calling, and with 

## TDDO
1. Data size as input;
2. Store status persistence and load;
3. Docker support