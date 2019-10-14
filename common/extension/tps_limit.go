/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package extension

import "github.com/apache/dubbo-go/filter"

var (
	tpsLimitStrategy = make(map[string]func(rate int, interval int) filter.TpsLimitStrategy)
	tpsLimiter = make(map[string]func() filter.TpsLimiter)
	tpsRejectedExecutionHandler = make(map[string]func() filter.RejectedExecutionHandler)
)

func SetTpsLimiter(name string, creator func() filter.TpsLimiter) {
	tpsLimiter[name] = creator
}

func GetTpsLimiter(name string) filter.TpsLimiter {
	var creator = tpsLimiter[name]
	if creator == nil {
		panic("TpsLimiter for " + name + " is not existing, make sure you have import the package " +
			"and you have register it by invoking extension.SetTpsLimiter.")
	}
	return creator()
}

func SetTpsLimitStrategy(name string, creator func(rate int, interval int) filter.TpsLimitStrategy) {
	tpsLimitStrategy[name] = creator
}

func GetTpsLimitStrategyCreator(name string) func(rate int, interval int) filter.TpsLimitStrategy {
	var creator = tpsLimitStrategy[name]
	if creator == nil {
		panic("TpsLimitStrategy for " + name + " is not existing, make sure you have import the package " +
			"and you have register it by invoking extension.SetTpsLimitStrategy.")
	}
	return creator
}

func SetTpsRejectedExecutionHandler(name string, creator func() filter.RejectedExecutionHandler)  {
	tpsRejectedExecutionHandler[name] = creator
}

func GetTpsRejectedExecutionHandler(name string) filter.RejectedExecutionHandler {
	var creator = tpsRejectedExecutionHandler[name]
	if creator() == nil {
		panic("TpsRejectedExecutionHandler for " + name + " is not existing, make sure you have import the package " +
			"and you have register it by invoking extension.SetTpsRejectedExecutionHandler.")
	}
	return creator()
}