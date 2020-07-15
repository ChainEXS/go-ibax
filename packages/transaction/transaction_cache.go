/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
	defer tc.mutex.RUnlock()

	t, ok = tc.cache[hash]
	return
}

func (tc *transactionCache) Set(t *Transaction) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tc.cache[string(t.TxHash)] = t
}

func (tc *transactionCache) Clean() {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tc.cache = make(map[string]*Transaction)
}