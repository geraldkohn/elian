// 使用useState创建Hook
import React, { useState } from 'react';

// 导入整个应用的Hook
import { useAppSelector, useAppDispatch } from '../../app/hooks';

// 导入自定义的Slice
import {
  decrement,
  increment,
  incrementByAmount,
  incrementAsync,
  incrementIfOdd,
  selectCount,
} from './counterSlice';

// 导入样式
import styles from './Counter.module.css';

export function Counter() {

  // 从store中获取state和dispatch
  const count = useAppSelector(selectCount);
  const dispatch = useAppDispatch();

  // 定义Hook的state
  const [incrementAmount, setIncrementAmount] = useState('2');
  // 定义increment值
  const incrementValue = Number(incrementAmount) || 0;

  return (
    <div>
      <div className={styles.row}>
        {/* 点击时 ()=>dispatch(function()) 分发 */}
        <button
          className={styles.button}
          aria-label="Decrement value"
          onClick={() => dispatch(decrement())}
        >
          -
        </button>

        <span className={styles.value}>{count}</span>

        <button
          className={styles.button}
          aria-label="Increment value"
          onClick={() => dispatch(increment())}
        >
          +
        </button>

      </div>


      <div className={styles.row}>
        {/* input改变时 event 中的target.value获取值 */}
        <input
          className={styles.textbox}
          aria-label="Set increment amount"
          value={incrementAmount}
          onChange={(e) => setIncrementAmount(e.target.value)}
        />
        
        <button
          className={styles.button}
          onClick={() => dispatch(incrementByAmount(incrementValue))}
        >
          Add Amount
        </button>
        
        <button
          className={styles.asyncButton}
          onClick={() => dispatch(incrementAsync(incrementValue))}
        >
          Add Async
        </button>
        
        <button
          className={styles.button}
          onClick={() => dispatch(incrementIfOdd(incrementValue))}
        >
          Add If Odd
        </button>

      </div>

    </div>
  );
}
