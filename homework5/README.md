# 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
~~~
./redis-benchmark -d 10 -t GET,SET                                                                                                                                        10:57:44
====== SET ======
  100000 requests completed in 1.65 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 1)
50.000% <= 0.423 milliseconds (cumulative count 53145)
75.000% <= 0.495 milliseconds (cumulative count 75067)
87.500% <= 0.599 milliseconds (cumulative count 87691)
93.750% <= 0.711 milliseconds (cumulative count 93979)
96.875% <= 0.791 milliseconds (cumulative count 96889)
98.438% <= 0.887 milliseconds (cumulative count 98493)
99.219% <= 1.023 milliseconds (cumulative count 99247)
99.609% <= 1.223 milliseconds (cumulative count 99613)
99.805% <= 1.583 milliseconds (cumulative count 99805)
99.902% <= 2.831 milliseconds (cumulative count 99904)
99.951% <= 3.239 milliseconds (cumulative count 99953)
99.976% <= 3.735 milliseconds (cumulative count 99976)
99.988% <= 5.311 milliseconds (cumulative count 99988)
99.994% <= 5.487 milliseconds (cumulative count 99994)
99.997% <= 5.623 milliseconds (cumulative count 99997)
99.998% <= 5.679 milliseconds (cumulative count 99999)
99.999% <= 5.719 milliseconds (cumulative count 100000)
100.000% <= 5.719 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.012% <= 0.207 milliseconds (cumulative count 12)
1.205% <= 0.303 milliseconds (cumulative count 1205)
43.958% <= 0.407 milliseconds (cumulative count 43958)
76.443% <= 0.503 milliseconds (cumulative count 76443)
88.303% <= 0.607 milliseconds (cumulative count 88303)
93.709% <= 0.703 milliseconds (cumulative count 93709)
97.346% <= 0.807 milliseconds (cumulative count 97346)
98.642% <= 0.903 milliseconds (cumulative count 98642)
99.189% <= 1.007 milliseconds (cumulative count 99189)
99.466% <= 1.103 milliseconds (cumulative count 99466)
99.604% <= 1.207 milliseconds (cumulative count 99604)
99.676% <= 1.303 milliseconds (cumulative count 99676)
99.757% <= 1.407 milliseconds (cumulative count 99757)
99.790% <= 1.503 milliseconds (cumulative count 99790)
99.808% <= 1.607 milliseconds (cumulative count 99808)
99.817% <= 1.703 milliseconds (cumulative count 99817)
99.822% <= 1.807 milliseconds (cumulative count 99822)
99.827% <= 1.903 milliseconds (cumulative count 99827)
99.833% <= 2.007 milliseconds (cumulative count 99833)
99.837% <= 2.103 milliseconds (cumulative count 99837)
99.936% <= 3.103 milliseconds (cumulative count 99936)
99.979% <= 4.103 milliseconds (cumulative count 99979)
100.000% <= 6.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 60642.81 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.462     0.144     0.423     0.743     0.959     5.719
====== GET ======
  100000 requests completed in 1.73 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.103 milliseconds (cumulative count 1)
50.000% <= 0.423 milliseconds (cumulative count 54050)
75.000% <= 0.495 milliseconds (cumulative count 75892)
87.500% <= 0.591 milliseconds (cumulative count 88079)
93.750% <= 0.679 milliseconds (cumulative count 94014)
96.875% <= 0.759 milliseconds (cumulative count 96962)
98.438% <= 0.839 milliseconds (cumulative count 98440)
99.219% <= 0.919 milliseconds (cumulative count 99240)
99.609% <= 0.999 milliseconds (cumulative count 99619)
99.805% <= 1.111 milliseconds (cumulative count 99805)
99.902% <= 1.303 milliseconds (cumulative count 99905)
99.951% <= 1.479 milliseconds (cumulative count 99952)
99.976% <= 4.535 milliseconds (cumulative count 99976)
99.988% <= 4.655 milliseconds (cumulative count 99990)
99.994% <= 4.727 milliseconds (cumulative count 99994)
99.997% <= 4.799 milliseconds (cumulative count 99998)
99.998% <= 4.807 milliseconds (cumulative count 99999)
99.999% <= 4.815 milliseconds (cumulative count 100000)
100.000% <= 4.815 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.001% <= 0.103 milliseconds (cumulative count 1)
0.013% <= 0.207 milliseconds (cumulative count 13)
0.667% <= 0.303 milliseconds (cumulative count 667)
40.938% <= 0.407 milliseconds (cumulative count 40938)
77.295% <= 0.503 milliseconds (cumulative count 77295)
89.447% <= 0.607 milliseconds (cumulative count 89447)
95.127% <= 0.703 milliseconds (cumulative count 95127)
97.992% <= 0.807 milliseconds (cumulative count 97992)
99.124% <= 0.903 milliseconds (cumulative count 99124)
99.644% <= 1.007 milliseconds (cumulative count 99644)
99.799% <= 1.103 milliseconds (cumulative count 99799)
99.861% <= 1.207 milliseconds (cumulative count 99861)
99.905% <= 1.303 milliseconds (cumulative count 99905)
99.937% <= 1.407 milliseconds (cumulative count 99937)
99.956% <= 1.503 milliseconds (cumulative count 99956)
99.964% <= 1.607 milliseconds (cumulative count 99964)
99.967% <= 1.703 milliseconds (cumulative count 99967)
99.969% <= 1.807 milliseconds (cumulative count 99969)
99.970% <= 1.903 milliseconds (cumulative count 99970)
99.972% <= 2.007 milliseconds (cumulative count 99972)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 57937.43 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.461     0.096     0.423     0.703     0.895     4.815


./redis-benchmark -d 20 -t GET,SET                                                                                                                                        10:58:39
====== SET ======
  100000 requests completed in 1.51 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 3)
50.000% <= 0.399 milliseconds (cumulative count 51868)
75.000% <= 0.447 milliseconds (cumulative count 76669)
87.500% <= 0.495 milliseconds (cumulative count 88398)
93.750% <= 0.551 milliseconds (cumulative count 94271)
96.875% <= 0.615 milliseconds (cumulative count 97022)
98.438% <= 0.695 milliseconds (cumulative count 98514)
99.219% <= 0.831 milliseconds (cumulative count 99236)
99.609% <= 1.039 milliseconds (cumulative count 99612)
99.805% <= 1.247 milliseconds (cumulative count 99807)
99.902% <= 1.447 milliseconds (cumulative count 99905)
99.951% <= 1.551 milliseconds (cumulative count 99953)
99.976% <= 1.687 milliseconds (cumulative count 99976)
99.988% <= 1.967 milliseconds (cumulative count 99988)
99.994% <= 2.167 milliseconds (cumulative count 99994)
99.997% <= 2.279 milliseconds (cumulative count 99997)
99.998% <= 2.367 milliseconds (cumulative count 99999)
99.999% <= 2.407 milliseconds (cumulative count 100000)
100.000% <= 2.407 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.009% <= 0.207 milliseconds (cumulative count 9)
0.598% <= 0.303 milliseconds (cumulative count 598)
57.266% <= 0.407 milliseconds (cumulative count 57266)
89.515% <= 0.503 milliseconds (cumulative count 89515)
96.803% <= 0.607 milliseconds (cumulative count 96803)
98.597% <= 0.703 milliseconds (cumulative count 98597)
99.153% <= 0.807 milliseconds (cumulative count 99153)
99.410% <= 0.903 milliseconds (cumulative count 99410)
99.576% <= 1.007 milliseconds (cumulative count 99576)
99.685% <= 1.103 milliseconds (cumulative count 99685)
99.779% <= 1.207 milliseconds (cumulative count 99779)
99.844% <= 1.303 milliseconds (cumulative count 99844)
99.885% <= 1.407 milliseconds (cumulative count 99885)
99.930% <= 1.503 milliseconds (cumulative count 99930)
99.969% <= 1.607 milliseconds (cumulative count 99969)
99.978% <= 1.703 milliseconds (cumulative count 99978)
99.982% <= 1.807 milliseconds (cumulative count 99982)
99.985% <= 1.903 milliseconds (cumulative count 99985)
99.989% <= 2.007 milliseconds (cumulative count 99989)
99.991% <= 2.103 milliseconds (cumulative count 99991)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 66269.05 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.419     0.144     0.399     0.567     0.767     2.407
====== GET ======
  100000 requests completed in 1.61 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 2)
50.000% <= 0.415 milliseconds (cumulative count 56573)
75.000% <= 0.447 milliseconds (cumulative count 78515)
87.500% <= 0.479 milliseconds (cumulative count 87534)
93.750% <= 0.535 milliseconds (cumulative count 93992)
96.875% <= 0.599 milliseconds (cumulative count 97042)
98.438% <= 0.663 milliseconds (cumulative count 98524)
99.219% <= 0.719 milliseconds (cumulative count 99244)
99.609% <= 0.783 milliseconds (cumulative count 99653)
99.805% <= 0.831 milliseconds (cumulative count 99821)
99.902% <= 0.879 milliseconds (cumulative count 99906)
99.951% <= 0.967 milliseconds (cumulative count 99954)
99.976% <= 1.023 milliseconds (cumulative count 99976)
99.988% <= 1.063 milliseconds (cumulative count 99992)
99.994% <= 1.071 milliseconds (cumulative count 99994)
99.997% <= 1.111 milliseconds (cumulative count 99997)
99.998% <= 1.151 milliseconds (cumulative count 99999)
99.999% <= 1.167 milliseconds (cumulative count 100000)
100.000% <= 1.167 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.009% <= 0.207 milliseconds (cumulative count 9)
0.076% <= 0.303 milliseconds (cumulative count 76)
47.733% <= 0.407 milliseconds (cumulative count 47733)
91.152% <= 0.503 milliseconds (cumulative count 91152)
97.302% <= 0.607 milliseconds (cumulative count 97302)
99.058% <= 0.703 milliseconds (cumulative count 99058)
99.743% <= 0.807 milliseconds (cumulative count 99743)
99.922% <= 0.903 milliseconds (cumulative count 99922)
99.967% <= 1.007 milliseconds (cumulative count 99967)
99.996% <= 1.103 milliseconds (cumulative count 99996)
100.000% <= 1.207 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 62305.30 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.427     0.136     0.415     0.559     0.703     1.167


./redis-benchmark -d 50 -t GET,SET                                                                                                                                        10:59:41
====== SET ======
  100000 requests completed in 1.98 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.127 milliseconds (cumulative count 1)
50.000% <= 0.447 milliseconds (cumulative count 50750)
75.000% <= 0.559 milliseconds (cumulative count 75750)
87.500% <= 0.671 milliseconds (cumulative count 87511)
93.750% <= 0.775 milliseconds (cumulative count 93933)
96.875% <= 0.895 milliseconds (cumulative count 96899)
98.438% <= 1.055 milliseconds (cumulative count 98469)
99.219% <= 1.271 milliseconds (cumulative count 99219)
99.609% <= 1.863 milliseconds (cumulative count 99611)
99.805% <= 2.935 milliseconds (cumulative count 99805)
99.902% <= 6.047 milliseconds (cumulative count 99904)
99.951% <= 13.935 milliseconds (cumulative count 99952)
99.976% <= 35.327 milliseconds (cumulative count 99976)
99.988% <= 51.327 milliseconds (cumulative count 99988)
99.994% <= 80.319 milliseconds (cumulative count 99994)
99.997% <= 80.383 milliseconds (cumulative count 99997)
99.998% <= 81.983 milliseconds (cumulative count 100000)
100.000% <= 81.983 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
0.120% <= 0.303 milliseconds (cumulative count 120)
26.631% <= 0.407 milliseconds (cumulative count 26631)
67.339% <= 0.503 milliseconds (cumulative count 67339)
81.357% <= 0.607 milliseconds (cumulative count 81357)
89.986% <= 0.703 milliseconds (cumulative count 89986)
95.047% <= 0.807 milliseconds (cumulative count 95047)
97.019% <= 0.903 milliseconds (cumulative count 97019)
98.113% <= 1.007 milliseconds (cumulative count 98113)
98.719% <= 1.103 milliseconds (cumulative count 98719)
99.098% <= 1.207 milliseconds (cumulative count 99098)
99.269% <= 1.303 milliseconds (cumulative count 99269)
99.381% <= 1.407 milliseconds (cumulative count 99381)
99.462% <= 1.503 milliseconds (cumulative count 99462)
99.517% <= 1.607 milliseconds (cumulative count 99517)
99.563% <= 1.703 milliseconds (cumulative count 99563)
99.602% <= 1.807 milliseconds (cumulative count 99602)
99.614% <= 1.903 milliseconds (cumulative count 99614)
99.628% <= 2.007 milliseconds (cumulative count 99628)
99.641% <= 2.103 milliseconds (cumulative count 99641)
99.808% <= 3.103 milliseconds (cumulative count 99808)
99.826% <= 4.103 milliseconds (cumulative count 99826)
99.866% <= 5.103 milliseconds (cumulative count 99866)
99.910% <= 6.103 milliseconds (cumulative count 99910)
99.922% <= 7.103 milliseconds (cumulative count 99922)
99.923% <= 9.103 milliseconds (cumulative count 99923)
99.925% <= 10.103 milliseconds (cumulative count 99925)
99.926% <= 12.103 milliseconds (cumulative count 99926)
99.957% <= 14.103 milliseconds (cumulative count 99957)
99.964% <= 15.103 milliseconds (cumulative count 99964)
99.977% <= 36.127 milliseconds (cumulative count 99977)
99.983% <= 51.103 milliseconds (cumulative count 99983)
99.990% <= 52.127 milliseconds (cumulative count 99990)
99.992% <= 80.127 milliseconds (cumulative count 99992)
99.998% <= 81.151 milliseconds (cumulative count 99998)
100.000% <= 82.111 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 50632.91 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.535     0.120     0.447     0.807     1.175    81.983
====== GET ======
  100000 requests completed in 2.09 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.207 milliseconds (cumulative count 1)
50.000% <= 0.471 milliseconds (cumulative count 52821)
75.000% <= 0.599 milliseconds (cumulative count 75834)
87.500% <= 0.735 milliseconds (cumulative count 87919)
93.750% <= 0.879 milliseconds (cumulative count 93894)
96.875% <= 1.063 milliseconds (cumulative count 96876)
98.438% <= 1.423 milliseconds (cumulative count 98443)
99.219% <= 1.967 milliseconds (cumulative count 99219)
99.609% <= 2.783 milliseconds (cumulative count 99610)
99.805% <= 4.151 milliseconds (cumulative count 99805)
99.902% <= 5.583 milliseconds (cumulative count 99905)
99.951% <= 6.207 milliseconds (cumulative count 99952)
99.976% <= 6.919 milliseconds (cumulative count 99976)
99.988% <= 9.623 milliseconds (cumulative count 99988)
99.994% <= 9.783 milliseconds (cumulative count 99994)
99.997% <= 9.855 milliseconds (cumulative count 99997)
99.998% <= 9.919 milliseconds (cumulative count 99999)
99.999% <= 9.943 milliseconds (cumulative count 100000)
100.000% <= 9.943 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.207 milliseconds (cumulative count 1)
0.104% <= 0.303 milliseconds (cumulative count 104)
15.814% <= 0.407 milliseconds (cumulative count 15814)
61.503% <= 0.503 milliseconds (cumulative count 61503)
76.654% <= 0.607 milliseconds (cumulative count 76654)
85.623% <= 0.703 milliseconds (cumulative count 85623)
91.715% <= 0.807 milliseconds (cumulative count 91715)
94.435% <= 0.903 milliseconds (cumulative count 94435)
96.232% <= 1.007 milliseconds (cumulative count 96232)
97.157% <= 1.103 milliseconds (cumulative count 97157)
97.747% <= 1.207 milliseconds (cumulative count 97747)
98.148% <= 1.303 milliseconds (cumulative count 98148)
98.406% <= 1.407 milliseconds (cumulative count 98406)
98.647% <= 1.503 milliseconds (cumulative count 98647)
98.789% <= 1.607 milliseconds (cumulative count 98789)
98.932% <= 1.703 milliseconds (cumulative count 98932)
99.071% <= 1.807 milliseconds (cumulative count 99071)
99.175% <= 1.903 milliseconds (cumulative count 99175)
99.246% <= 2.007 milliseconds (cumulative count 99246)
99.316% <= 2.103 milliseconds (cumulative count 99316)
99.691% <= 3.103 milliseconds (cumulative count 99691)
99.801% <= 4.103 milliseconds (cumulative count 99801)
99.882% <= 5.103 milliseconds (cumulative count 99882)
99.947% <= 6.103 milliseconds (cumulative count 99947)
99.976% <= 7.103 milliseconds (cumulative count 99976)
100.000% <= 10.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 47869.79 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.558     0.200     0.471     0.935     1.759     9.943

./redis-benchmark -d 50 -t GET,SET                                                                                                                                        10:59:41
====== SET ======
  100000 requests completed in 1.98 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.127 milliseconds (cumulative count 1)
50.000% <= 0.447 milliseconds (cumulative count 50750)
75.000% <= 0.559 milliseconds (cumulative count 75750)
87.500% <= 0.671 milliseconds (cumulative count 87511)
93.750% <= 0.775 milliseconds (cumulative count 93933)
96.875% <= 0.895 milliseconds (cumulative count 96899)
98.438% <= 1.055 milliseconds (cumulative count 98469)
99.219% <= 1.271 milliseconds (cumulative count 99219)
99.609% <= 1.863 milliseconds (cumulative count 99611)
99.805% <= 2.935 milliseconds (cumulative count 99805)
99.902% <= 6.047 milliseconds (cumulative count 99904)
99.951% <= 13.935 milliseconds (cumulative count 99952)
99.976% <= 35.327 milliseconds (cumulative count 99976)
99.988% <= 51.327 milliseconds (cumulative count 99988)
99.994% <= 80.319 milliseconds (cumulative count 99994)
99.997% <= 80.383 milliseconds (cumulative count 99997)
99.998% <= 81.983 milliseconds (cumulative count 100000)
100.000% <= 81.983 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
0.120% <= 0.303 milliseconds (cumulative count 120)
26.631% <= 0.407 milliseconds (cumulative count 26631)
67.339% <= 0.503 milliseconds (cumulative count 67339)
81.357% <= 0.607 milliseconds (cumulative count 81357)
89.986% <= 0.703 milliseconds (cumulative count 89986)
95.047% <= 0.807 milliseconds (cumulative count 95047)
97.019% <= 0.903 milliseconds (cumulative count 97019)
98.113% <= 1.007 milliseconds (cumulative count 98113)
98.719% <= 1.103 milliseconds (cumulative count 98719)
99.098% <= 1.207 milliseconds (cumulative count 99098)
99.269% <= 1.303 milliseconds (cumulative count 99269)
99.381% <= 1.407 milliseconds (cumulative count 99381)
99.462% <= 1.503 milliseconds (cumulative count 99462)
99.517% <= 1.607 milliseconds (cumulative count 99517)
99.563% <= 1.703 milliseconds (cumulative count 99563)
99.602% <= 1.807 milliseconds (cumulative count 99602)
99.614% <= 1.903 milliseconds (cumulative count 99614)
99.628% <= 2.007 milliseconds (cumulative count 99628)
99.641% <= 2.103 milliseconds (cumulative count 99641)
99.808% <= 3.103 milliseconds (cumulative count 99808)
99.826% <= 4.103 milliseconds (cumulative count 99826)
99.866% <= 5.103 milliseconds (cumulative count 99866)
99.910% <= 6.103 milliseconds (cumulative count 99910)
99.922% <= 7.103 milliseconds (cumulative count 99922)
99.923% <= 9.103 milliseconds (cumulative count 99923)
99.925% <= 10.103 milliseconds (cumulative count 99925)
99.926% <= 12.103 milliseconds (cumulative count 99926)
99.957% <= 14.103 milliseconds (cumulative count 99957)
99.964% <= 15.103 milliseconds (cumulative count 99964)
99.977% <= 36.127 milliseconds (cumulative count 99977)
99.983% <= 51.103 milliseconds (cumulative count 99983)
99.990% <= 52.127 milliseconds (cumulative count 99990)
99.992% <= 80.127 milliseconds (cumulative count 99992)
99.998% <= 81.151 milliseconds (cumulative count 99998)
100.000% <= 82.111 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 50632.91 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.535     0.120     0.447     0.807     1.175    81.983
====== GET ======
  100000 requests completed in 2.09 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.207 milliseconds (cumulative count 1)
50.000% <= 0.471 milliseconds (cumulative count 52821)
75.000% <= 0.599 milliseconds (cumulative count 75834)
87.500% <= 0.735 milliseconds (cumulative count 87919)
93.750% <= 0.879 milliseconds (cumulative count 93894)
96.875% <= 1.063 milliseconds (cumulative count 96876)
98.438% <= 1.423 milliseconds (cumulative count 98443)
99.219% <= 1.967 milliseconds (cumulative count 99219)
99.609% <= 2.783 milliseconds (cumulative count 99610)
99.805% <= 4.151 milliseconds (cumulative count 99805)
99.902% <= 5.583 milliseconds (cumulative count 99905)
99.951% <= 6.207 milliseconds (cumulative count 99952)
99.976% <= 6.919 milliseconds (cumulative count 99976)
99.988% <= 9.623 milliseconds (cumulative count 99988)
99.994% <= 9.783 milliseconds (cumulative count 99994)
99.997% <= 9.855 milliseconds (cumulative count 99997)
99.998% <= 9.919 milliseconds (cumulative count 99999)
99.999% <= 9.943 milliseconds (cumulative count 100000)
100.000% <= 9.943 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.207 milliseconds (cumulative count 1)
0.104% <= 0.303 milliseconds (cumulative count 104)
15.814% <= 0.407 milliseconds (cumulative count 15814)
61.503% <= 0.503 milliseconds (cumulative count 61503)
76.654% <= 0.607 milliseconds (cumulative count 76654)
85.623% <= 0.703 milliseconds (cumulative count 85623)
91.715% <= 0.807 milliseconds (cumulative count 91715)
94.435% <= 0.903 milliseconds (cumulative count 94435)
96.232% <= 1.007 milliseconds (cumulative count 96232)
97.157% <= 1.103 milliseconds (cumulative count 97157)
97.747% <= 1.207 milliseconds (cumulative count 97747)
98.148% <= 1.303 milliseconds (cumulative count 98148)
98.406% <= 1.407 milliseconds (cumulative count 98406)
98.647% <= 1.503 milliseconds (cumulative count 98647)
98.789% <= 1.607 milliseconds (cumulative count 98789)
98.932% <= 1.703 milliseconds (cumulative count 98932)
99.071% <= 1.807 milliseconds (cumulative count 99071)
99.175% <= 1.903 milliseconds (cumulative count 99175)
99.246% <= 2.007 milliseconds (cumulative count 99246)
99.316% <= 2.103 milliseconds (cumulative count 99316)
99.691% <= 3.103 milliseconds (cumulative count 99691)
99.801% <= 4.103 milliseconds (cumulative count 99801)
99.882% <= 5.103 milliseconds (cumulative count 99882)
99.947% <= 6.103 milliseconds (cumulative count 99947)
99.976% <= 7.103 milliseconds (cumulative count 99976)
100.000% <= 10.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 47869.79 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.558     0.200     0.471     0.935     1.759     9.943

./redis-benchmark -d 100 -t GET,SET                                                                                                                                       11:00:11
====== SET ======
  100000 requests completed in 1.69 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.447 milliseconds (cumulative count 51103)
75.000% <= 0.543 milliseconds (cumulative count 75066)
87.500% <= 0.639 milliseconds (cumulative count 88029)
93.750% <= 0.727 milliseconds (cumulative count 93788)
96.875% <= 0.823 milliseconds (cumulative count 96959)
98.438% <= 0.935 milliseconds (cumulative count 98511)
99.219% <= 1.047 milliseconds (cumulative count 99237)
99.609% <= 1.231 milliseconds (cumulative count 99618)
99.805% <= 1.495 milliseconds (cumulative count 99807)
99.902% <= 2.223 milliseconds (cumulative count 99903)
99.951% <= 2.767 milliseconds (cumulative count 99952)
99.976% <= 2.975 milliseconds (cumulative count 99976)
99.988% <= 3.223 milliseconds (cumulative count 99988)
99.994% <= 3.439 milliseconds (cumulative count 99994)
99.997% <= 3.527 milliseconds (cumulative count 99997)
99.998% <= 3.607 milliseconds (cumulative count 99999)
99.999% <= 3.639 milliseconds (cumulative count 100000)
100.000% <= 3.639 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
0.508% <= 0.303 milliseconds (cumulative count 508)
34.373% <= 0.407 milliseconds (cumulative count 34373)
67.321% <= 0.503 milliseconds (cumulative count 67321)
84.522% <= 0.607 milliseconds (cumulative count 84522)
92.604% <= 0.703 milliseconds (cumulative count 92604)
96.580% <= 0.807 milliseconds (cumulative count 96580)
98.139% <= 0.903 milliseconds (cumulative count 98139)
99.059% <= 1.007 milliseconds (cumulative count 99059)
99.368% <= 1.103 milliseconds (cumulative count 99368)
99.574% <= 1.207 milliseconds (cumulative count 99574)
99.705% <= 1.303 milliseconds (cumulative count 99705)
99.783% <= 1.407 milliseconds (cumulative count 99783)
99.808% <= 1.503 milliseconds (cumulative count 99808)
99.823% <= 1.607 milliseconds (cumulative count 99823)
99.834% <= 1.703 milliseconds (cumulative count 99834)
99.848% <= 1.807 milliseconds (cumulative count 99848)
99.861% <= 1.903 milliseconds (cumulative count 99861)
99.880% <= 2.007 milliseconds (cumulative count 99880)
99.891% <= 2.103 milliseconds (cumulative count 99891)
99.983% <= 3.103 milliseconds (cumulative count 99983)
100.000% <= 4.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 59241.71 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.487     0.112     0.447     0.759     0.999     3.639
====== GET ======
  100000 requests completed in 1.74 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.103 milliseconds (cumulative count 1)
50.000% <= 0.431 milliseconds (cumulative count 51139)
75.000% <= 0.527 milliseconds (cumulative count 75353)
87.500% <= 0.631 milliseconds (cumulative count 87658)
93.750% <= 0.719 milliseconds (cumulative count 93981)
96.875% <= 0.791 milliseconds (cumulative count 96898)
98.438% <= 0.871 milliseconds (cumulative count 98439)
99.219% <= 0.943 milliseconds (cumulative count 99245)
99.609% <= 1.015 milliseconds (cumulative count 99631)
99.805% <= 1.087 milliseconds (cumulative count 99822)
99.902% <= 1.175 milliseconds (cumulative count 99904)
99.951% <= 1.335 milliseconds (cumulative count 99952)
99.976% <= 1.423 milliseconds (cumulative count 99976)
99.988% <= 1.479 milliseconds (cumulative count 99988)
99.994% <= 1.519 milliseconds (cumulative count 99994)
99.997% <= 1.559 milliseconds (cumulative count 99998)
99.998% <= 1.567 milliseconds (cumulative count 99999)
99.999% <= 1.639 milliseconds (cumulative count 100000)
100.000% <= 1.639 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.001% <= 0.103 milliseconds (cumulative count 1)
0.012% <= 0.207 milliseconds (cumulative count 12)
0.324% <= 0.303 milliseconds (cumulative count 324)
35.221% <= 0.407 milliseconds (cumulative count 35221)
71.554% <= 0.503 milliseconds (cumulative count 71554)
85.166% <= 0.607 milliseconds (cumulative count 85166)
93.052% <= 0.703 milliseconds (cumulative count 93052)
97.304% <= 0.807 milliseconds (cumulative count 97304)
98.873% <= 0.903 milliseconds (cumulative count 98873)
99.595% <= 1.007 milliseconds (cumulative count 99595)
99.843% <= 1.103 milliseconds (cumulative count 99843)
99.917% <= 1.207 milliseconds (cumulative count 99917)
99.944% <= 1.303 milliseconds (cumulative count 99944)
99.972% <= 1.407 milliseconds (cumulative count 99972)
99.990% <= 1.503 milliseconds (cumulative count 99990)
99.999% <= 1.607 milliseconds (cumulative count 99999)
100.000% <= 1.703 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 57636.89 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.477     0.096     0.431     0.743     0.919     1.639

./redis-benchmark -d 200 -t GET,SET                                                                                                                                       11:02:12
====== SET ======
  100000 requests completed in 1.61 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.439 milliseconds (cumulative count 50239)
75.000% <= 0.535 milliseconds (cumulative count 76314)
87.500% <= 0.615 milliseconds (cumulative count 88057)
93.750% <= 0.695 milliseconds (cumulative count 93774)
96.875% <= 0.783 milliseconds (cumulative count 96956)
98.438% <= 0.879 milliseconds (cumulative count 98459)
99.219% <= 0.991 milliseconds (cumulative count 99233)
99.609% <= 1.175 milliseconds (cumulative count 99617)
99.805% <= 1.423 milliseconds (cumulative count 99806)
99.902% <= 2.431 milliseconds (cumulative count 99906)
99.951% <= 2.551 milliseconds (cumulative count 99952)
99.976% <= 2.743 milliseconds (cumulative count 99976)
99.988% <= 3.055 milliseconds (cumulative count 99988)
99.994% <= 3.255 milliseconds (cumulative count 99994)
99.997% <= 3.527 milliseconds (cumulative count 99997)
99.998% <= 3.591 milliseconds (cumulative count 99999)
99.999% <= 3.639 milliseconds (cumulative count 100000)
100.000% <= 3.639 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.002% <= 0.207 milliseconds (cumulative count 2)
1.241% <= 0.303 milliseconds (cumulative count 1241)
37.075% <= 0.407 milliseconds (cumulative count 37075)
69.502% <= 0.503 milliseconds (cumulative count 69502)
87.158% <= 0.607 milliseconds (cumulative count 87158)
94.172% <= 0.703 milliseconds (cumulative count 94172)
97.488% <= 0.807 milliseconds (cumulative count 97488)
98.704% <= 0.903 milliseconds (cumulative count 98704)
99.287% <= 1.007 milliseconds (cumulative count 99287)
99.516% <= 1.103 milliseconds (cumulative count 99516)
99.658% <= 1.207 milliseconds (cumulative count 99658)
99.757% <= 1.303 milliseconds (cumulative count 99757)
99.801% <= 1.407 milliseconds (cumulative count 99801)
99.825% <= 1.503 milliseconds (cumulative count 99825)
99.848% <= 1.607 milliseconds (cumulative count 99848)
99.861% <= 1.703 milliseconds (cumulative count 99861)
99.866% <= 1.807 milliseconds (cumulative count 99866)
99.869% <= 1.903 milliseconds (cumulative count 99869)
99.874% <= 2.007 milliseconds (cumulative count 99874)
99.879% <= 2.103 milliseconds (cumulative count 99879)
99.989% <= 3.103 milliseconds (cumulative count 99989)
100.000% <= 4.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 61996.28 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.474     0.128     0.439     0.727     0.951     3.639
====== GET ======
  100000 requests completed in 1.77 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.095 milliseconds (cumulative count 1)
50.000% <= 0.431 milliseconds (cumulative count 51757)
75.000% <= 0.535 milliseconds (cumulative count 75334)
87.500% <= 0.647 milliseconds (cumulative count 88220)
93.750% <= 0.743 milliseconds (cumulative count 94090)
96.875% <= 0.839 milliseconds (cumulative count 96915)
98.438% <= 0.943 milliseconds (cumulative count 98461)
99.219% <= 1.039 milliseconds (cumulative count 99242)
99.609% <= 1.159 milliseconds (cumulative count 99621)
99.805% <= 1.303 milliseconds (cumulative count 99808)
99.902% <= 1.399 milliseconds (cumulative count 99920)
99.951% <= 1.423 milliseconds (cumulative count 99959)
99.976% <= 1.519 milliseconds (cumulative count 99976)
99.988% <= 1.871 milliseconds (cumulative count 99988)
99.994% <= 1.903 milliseconds (cumulative count 99995)
99.997% <= 1.927 milliseconds (cumulative count 99997)
99.998% <= 1.983 milliseconds (cumulative count 99999)
99.999% <= 2.031 milliseconds (cumulative count 100000)
100.000% <= 2.031 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.001% <= 0.103 milliseconds (cumulative count 1)
0.009% <= 0.207 milliseconds (cumulative count 9)
0.358% <= 0.303 milliseconds (cumulative count 358)
37.453% <= 0.407 milliseconds (cumulative count 37453)
70.361% <= 0.503 milliseconds (cumulative count 70361)
84.371% <= 0.607 milliseconds (cumulative count 84371)
92.138% <= 0.703 milliseconds (cumulative count 92138)
96.212% <= 0.807 milliseconds (cumulative count 96212)
98.004% <= 0.903 milliseconds (cumulative count 98004)
99.014% <= 1.007 milliseconds (cumulative count 99014)
99.499% <= 1.103 milliseconds (cumulative count 99499)
99.705% <= 1.207 milliseconds (cumulative count 99705)
99.808% <= 1.303 milliseconds (cumulative count 99808)
99.937% <= 1.407 milliseconds (cumulative count 99937)
99.974% <= 1.503 milliseconds (cumulative count 99974)
99.977% <= 1.607 milliseconds (cumulative count 99977)
99.979% <= 1.703 milliseconds (cumulative count 99979)
99.981% <= 1.807 milliseconds (cumulative count 99981)
99.995% <= 1.903 milliseconds (cumulative count 99995)
99.999% <= 2.007 milliseconds (cumulative count 99999)
100.000% <= 2.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 56529.11 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.480     0.088     0.431     0.767     1.007     2.031

./redis-benchmark -d 1000 -t GET,SET                                                                                                                                      11:03:19
====== SET ======
  100000 requests completed in 1.61 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.431 milliseconds (cumulative count 52662)
75.000% <= 0.519 milliseconds (cumulative count 75524)
87.500% <= 0.599 milliseconds (cumulative count 87611)
93.750% <= 0.671 milliseconds (cumulative count 93751)
96.875% <= 0.767 milliseconds (cumulative count 96958)
98.438% <= 0.871 milliseconds (cumulative count 98518)
99.219% <= 0.959 milliseconds (cumulative count 99237)
99.609% <= 1.063 milliseconds (cumulative count 99612)
99.805% <= 1.167 milliseconds (cumulative count 99806)
99.902% <= 1.263 milliseconds (cumulative count 99909)
99.951% <= 1.431 milliseconds (cumulative count 99953)
99.976% <= 1.791 milliseconds (cumulative count 99976)
99.988% <= 1.951 milliseconds (cumulative count 99990)
99.994% <= 1.999 milliseconds (cumulative count 99994)
99.997% <= 2.055 milliseconds (cumulative count 99997)
99.998% <= 2.095 milliseconds (cumulative count 99999)
99.999% <= 2.119 milliseconds (cumulative count 100000)
100.000% <= 2.119 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.009% <= 0.207 milliseconds (cumulative count 9)
1.402% <= 0.303 milliseconds (cumulative count 1402)
43.213% <= 0.407 milliseconds (cumulative count 43213)
72.479% <= 0.503 milliseconds (cumulative count 72479)
88.557% <= 0.607 milliseconds (cumulative count 88557)
95.193% <= 0.703 milliseconds (cumulative count 95193)
97.683% <= 0.807 milliseconds (cumulative count 97683)
98.827% <= 0.903 milliseconds (cumulative count 98827)
99.445% <= 1.007 milliseconds (cumulative count 99445)
99.698% <= 1.103 milliseconds (cumulative count 99698)
99.857% <= 1.207 milliseconds (cumulative count 99857)
99.928% <= 1.303 milliseconds (cumulative count 99928)
99.949% <= 1.407 milliseconds (cumulative count 99949)
99.958% <= 1.503 milliseconds (cumulative count 99958)
99.965% <= 1.607 milliseconds (cumulative count 99965)
99.970% <= 1.703 milliseconds (cumulative count 99970)
99.977% <= 1.807 milliseconds (cumulative count 99977)
99.983% <= 1.903 milliseconds (cumulative count 99983)
99.994% <= 2.007 milliseconds (cumulative count 99994)
99.999% <= 2.103 milliseconds (cumulative count 99999)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 62034.74 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.460     0.112     0.431     0.703     0.927     2.119
====== GET ======
  100000 requests completed in 1.80 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 2)
50.000% <= 0.447 milliseconds (cumulative count 50756)
75.000% <= 0.567 milliseconds (cumulative count 75833)
87.500% <= 0.655 milliseconds (cumulative count 88064)
93.750% <= 0.735 milliseconds (cumulative count 94044)
96.875% <= 0.815 milliseconds (cumulative count 96897)
98.438% <= 0.895 milliseconds (cumulative count 98541)
99.219% <= 0.967 milliseconds (cumulative count 99223)
99.609% <= 1.087 milliseconds (cumulative count 99611)
99.805% <= 1.239 milliseconds (cumulative count 99808)
99.902% <= 1.431 milliseconds (cumulative count 99906)
99.951% <= 1.607 milliseconds (cumulative count 99952)
99.976% <= 1.815 milliseconds (cumulative count 99976)
99.988% <= 2.063 milliseconds (cumulative count 99988)
99.994% <= 2.191 milliseconds (cumulative count 99994)
99.997% <= 2.295 milliseconds (cumulative count 99997)
99.998% <= 2.335 milliseconds (cumulative count 99999)
99.999% <= 2.415 milliseconds (cumulative count 100000)
100.000% <= 2.415 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 0.207 milliseconds (cumulative count 4)
0.067% <= 0.303 milliseconds (cumulative count 67)
25.653% <= 0.407 milliseconds (cumulative count 25653)
65.298% <= 0.503 milliseconds (cumulative count 65298)
81.743% <= 0.607 milliseconds (cumulative count 81743)
92.173% <= 0.703 milliseconds (cumulative count 92173)
96.700% <= 0.807 milliseconds (cumulative count 96700)
98.656% <= 0.903 milliseconds (cumulative count 98656)
99.402% <= 1.007 milliseconds (cumulative count 99402)
99.630% <= 1.103 milliseconds (cumulative count 99630)
99.766% <= 1.207 milliseconds (cumulative count 99766)
99.857% <= 1.303 milliseconds (cumulative count 99857)
99.896% <= 1.407 milliseconds (cumulative count 99896)
99.927% <= 1.503 milliseconds (cumulative count 99927)
99.952% <= 1.607 milliseconds (cumulative count 99952)
99.964% <= 1.703 milliseconds (cumulative count 99964)
99.975% <= 1.807 milliseconds (cumulative count 99975)
99.984% <= 1.903 milliseconds (cumulative count 99984)
99.987% <= 2.007 milliseconds (cumulative count 99987)
99.988% <= 2.103 milliseconds (cumulative count 99988)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 55401.66 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.497     0.144     0.447     0.759     0.943     2.415


src ./redis-benchmark -d 5000 -t GET,SET                                                                                                                                      11:04:47
====== SET ======
  100000 requests completed in 1.65 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.167 milliseconds (cumulative count 1)
50.000% <= 0.439 milliseconds (cumulative count 52152)
75.000% <= 0.511 milliseconds (cumulative count 75634)
87.500% <= 0.607 milliseconds (cumulative count 87854)
93.750% <= 0.703 milliseconds (cumulative count 94009)
96.875% <= 0.815 milliseconds (cumulative count 96937)
98.438% <= 0.983 milliseconds (cumulative count 98465)
99.219% <= 1.191 milliseconds (cumulative count 99230)
99.609% <= 1.471 milliseconds (cumulative count 99616)
99.805% <= 1.903 milliseconds (cumulative count 99824)
99.902% <= 1.951 milliseconds (cumulative count 99908)
99.951% <= 2.111 milliseconds (cumulative count 99952)
99.976% <= 2.479 milliseconds (cumulative count 99976)
99.988% <= 2.815 milliseconds (cumulative count 99988)
99.994% <= 2.935 milliseconds (cumulative count 99994)
99.997% <= 3.015 milliseconds (cumulative count 99998)
99.998% <= 3.039 milliseconds (cumulative count 99999)
99.999% <= 3.071 milliseconds (cumulative count 100000)
100.000% <= 3.071 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.002% <= 0.207 milliseconds (cumulative count 2)
2.733% <= 0.303 milliseconds (cumulative count 2733)
35.824% <= 0.407 milliseconds (cumulative count 35824)
73.906% <= 0.503 milliseconds (cumulative count 73906)
87.854% <= 0.607 milliseconds (cumulative count 87854)
94.009% <= 0.703 milliseconds (cumulative count 94009)
96.806% <= 0.807 milliseconds (cumulative count 96806)
97.934% <= 0.903 milliseconds (cumulative count 97934)
98.604% <= 1.007 milliseconds (cumulative count 98604)
98.972% <= 1.103 milliseconds (cumulative count 98972)
99.268% <= 1.207 milliseconds (cumulative count 99268)
99.477% <= 1.303 milliseconds (cumulative count 99477)
99.561% <= 1.407 milliseconds (cumulative count 99561)
99.647% <= 1.503 milliseconds (cumulative count 99647)
99.717% <= 1.607 milliseconds (cumulative count 99717)
99.749% <= 1.703 milliseconds (cumulative count 99749)
99.770% <= 1.807 milliseconds (cumulative count 99770)
99.824% <= 1.903 milliseconds (cumulative count 99824)
99.944% <= 2.007 milliseconds (cumulative count 99944)
99.950% <= 2.103 milliseconds (cumulative count 99950)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 60679.61 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.472     0.160     0.439     0.735     1.119     3.071
====== GET ======
  100000 requests completed in 1.81 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.439 milliseconds (cumulative count 54541)
75.000% <= 0.519 milliseconds (cumulative count 75475)
87.500% <= 0.623 milliseconds (cumulative count 87801)
93.750% <= 0.711 milliseconds (cumulative count 93812)
96.875% <= 0.807 milliseconds (cumulative count 96942)
98.438% <= 0.919 milliseconds (cumulative count 98443)
99.219% <= 1.055 milliseconds (cumulative count 99231)
99.609% <= 1.271 milliseconds (cumulative count 99610)
99.805% <= 1.487 milliseconds (cumulative count 99808)
99.902% <= 1.695 milliseconds (cumulative count 99904)
99.951% <= 1.847 milliseconds (cumulative count 99952)
99.976% <= 2.071 milliseconds (cumulative count 99976)
99.988% <= 2.303 milliseconds (cumulative count 99988)
99.994% <= 2.527 milliseconds (cumulative count 99994)
99.997% <= 2.831 milliseconds (cumulative count 99997)
99.998% <= 3.383 milliseconds (cumulative count 99999)
99.999% <= 3.471 milliseconds (cumulative count 100000)
100.000% <= 3.471 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
0.014% <= 0.303 milliseconds (cumulative count 14)
26.767% <= 0.407 milliseconds (cumulative count 26767)
72.681% <= 0.503 milliseconds (cumulative count 72681)
86.263% <= 0.607 milliseconds (cumulative count 86263)
93.432% <= 0.703 milliseconds (cumulative count 93432)
96.942% <= 0.807 milliseconds (cumulative count 96942)
98.320% <= 0.903 milliseconds (cumulative count 98320)
99.042% <= 1.007 milliseconds (cumulative count 99042)
99.351% <= 1.103 milliseconds (cumulative count 99351)
99.515% <= 1.207 milliseconds (cumulative count 99515)
99.646% <= 1.303 milliseconds (cumulative count 99646)
99.747% <= 1.407 milliseconds (cumulative count 99747)
99.813% <= 1.503 milliseconds (cumulative count 99813)
99.864% <= 1.607 milliseconds (cumulative count 99864)
99.907% <= 1.703 milliseconds (cumulative count 99907)
99.942% <= 1.807 milliseconds (cumulative count 99942)
99.961% <= 1.903 milliseconds (cumulative count 99961)
99.973% <= 2.007 milliseconds (cumulative count 99973)
99.977% <= 2.103 milliseconds (cumulative count 99977)
99.998% <= 3.103 milliseconds (cumulative count 99998)
100.000% <= 4.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 55340.34 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.485     0.128     0.439     0.743     1.007     3.471
~~~

# 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
~~~
info memory
# Memory
used_memory:2288288
used_memory_human:2.18M
used_memory_rss:3706880
used_memory_rss_human:3.54M
used_memory_peak:2348768
used_memory_peak_human:2.24M
used_memory_peak_perc:97.43%
used_memory_overhead:1029440
used_memory_startup:1011776
used_memory_dataset:1258848
used_memory_dataset_perc:98.62%
allocator_allocated:2243168
allocator_active:3668992
allocator_resident:3668992
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.64
allocator_frag_bytes:1425824
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.65
mem_fragmentation_bytes:1463712
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

./redis-benchmark -t set -d 50000 -r 10000000 -n 10000                                                                                                                    12:02:54
====== SET ======
  10000 requests completed in 0.63 seconds
  50 parallel clients
  50000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.119 milliseconds (cumulative count 1)
50.000% <= 2.743 milliseconds (cumulative count 5001)
75.000% <= 3.335 milliseconds (cumulative count 7508)
87.500% <= 4.031 milliseconds (cumulative count 8759)
93.750% <= 4.559 milliseconds (cumulative count 9380)
96.875% <= 5.447 milliseconds (cumulative count 9690)
98.438% <= 6.359 milliseconds (cumulative count 9844)
99.219% <= 7.055 milliseconds (cumulative count 9922)
99.609% <= 7.543 milliseconds (cumulative count 9961)
99.805% <= 7.903 milliseconds (cumulative count 9981)
99.902% <= 8.095 milliseconds (cumulative count 9992)
99.951% <= 8.303 milliseconds (cumulative count 9996)
99.976% <= 8.407 milliseconds (cumulative count 9998)
99.988% <= 8.431 milliseconds (cumulative count 9999)
99.994% <= 8.519 milliseconds (cumulative count 10000)
100.000% <= 8.519 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.020% <= 1.207 milliseconds (cumulative count 2)
0.030% <= 1.407 milliseconds (cumulative count 3)
0.060% <= 1.503 milliseconds (cumulative count 6)
0.100% <= 1.607 milliseconds (cumulative count 10)
0.140% <= 1.703 milliseconds (cumulative count 14)
0.300% <= 1.807 milliseconds (cumulative count 30)
1.490% <= 1.903 milliseconds (cumulative count 149)
5.800% <= 2.007 milliseconds (cumulative count 580)
12.130% <= 2.103 milliseconds (cumulative count 1213)
66.600% <= 3.103 milliseconds (cumulative count 6660)
88.510% <= 4.103 milliseconds (cumulative count 8851)
95.920% <= 5.103 milliseconds (cumulative count 9592)
98.110% <= 6.103 milliseconds (cumulative count 9811)
99.290% <= 7.103 milliseconds (cumulative count 9929)
99.920% <= 8.103 milliseconds (cumulative count 9992)
100.000% <= 9.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 15898.25 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.978     1.112     2.743     4.799     6.831     8.519


./redis-cli                                                                                                                                                               12:14:47
127.0.0.1:6379> info memory
# Memory
used_memory:537709744
used_memory_human:512.80M
used_memory_rss:536129536
used_memory_rss_human:511.29M
used_memory_peak:541196736
used_memory_peak_human:516.13M
used_memory_peak_perc:99.36%
used_memory_overhead:1560120
used_memory_startup:1011776
used_memory_dataset:536149624
used_memory_dataset_perc:99.90%
allocator_allocated:537656432
allocator_active:536091648
allocator_resident:536091648
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:18446744073707986832
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.00
mem_fragmentation_bytes:-1526896
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
~~~

avg_key_size = 537709744-2288288-500000000=35421456/10000=3542