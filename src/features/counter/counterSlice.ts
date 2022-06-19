import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { fetchCount } from './counterAPI';
import { RootState, AppThunk } from '../../app/store';

// 导出接口  值为number类型  状态为 idle loading failed三个字符串
export interface CounterState {
  value: number;
  status: 'idle' | 'loading' | 'failed';
}

// 定义初始化状态  类型为上方定义的接口  counter 值为0 状态为idle
const initialState: CounterState = {
  value: 0,
  status: 'idle',
};



// 导入我们自定义的异步方法

// 定义了一个异步执行的逻辑  可以像一个普通的action一样被dispatch  这会变为第一个参数
// 异步代码可以之后被执行 其他的action也可以被dispatch
export const incrementAsync = createAsyncThunk(
  'counter/fetchCount',
  async (amount: number) => {
    const response = await fetchCount(amount);
    // 我们return的值 将会变成'fulfilled' action payload
    return response.data;
  }
);

// 导出一个Slice
export const counterSlice = createSlice({
  name: 'counter',
  initialState,
  // reduer字段允许我们定义reducers和生成相关的action
  reducers: {
    increment: (state) => {
      // Redux Toolkit允许我们写一个变化逻辑  事实上不用改变state 因为它使用immer库(探测draft state的改变 基于这些变化生成一个全新的不可变state)
      state.value += 1;
    },
    decrement: (state) => {
      state.value -= 1;
    },
    // 使用payloadAction类型来声明action.payload的内容
    incrementByAmount: (state, action: PayloadAction<number>) => {
      state.value += action.payload;
    },
  },
  // extraReducers字段允许这个slice操作定义在别处的action  包括异步的action或者定义在其他slice的action
  // 改变status的值 通过异步请求
  extraReducers: (builder) => {
    builder
      .addCase(incrementAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(incrementAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.value += action.payload;
      })
      .addCase(incrementAsync.rejected, (state) => {
        state.status = 'failed';
      });
  },
});

// 导出这个Slice中的Action
export const { increment, decrement, incrementByAmount } = counterSlice.actions;



// 导入自定义的redux类型用于ts规范


// selector 允许我们从状态中取值  Selector也可以在使用的时候inline定义 而不是在slice文件中  例如`useSelector((state: RootState) => state.counter.value)`
export const selectCount = (state: RootState) => state.counter.value;


// 我们也可以手动编写可能同时包含同步和异步逻辑的thunk  下面是一个基于条件而选择性改变状态的案例
export const incrementIfOdd =
  (amount: number): AppThunk =>
  (dispatch, getState) => {
    const currentValue = selectCount(getState());
    if (currentValue % 2 === 1) {
      dispatch(incrementByAmount(amount));
    }
  };

// default导出的是Slice的Reducer
export default counterSlice.reducer;
