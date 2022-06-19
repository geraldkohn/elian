// 引入configureStore创建全局store
import { configureStore} from '@reduxjs/toolkit';
import { ThunkAction, Action } from '@reduxjs/toolkit';
// 从自定一的Slice引入自定义的Reducer
import counterReducer from '../features/counter/counterSlice';
export const store = configureStore({
  reducer: {
    counter: counterReducer,
  },
});


// 定义dispatch和state类型
export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;


// 引入Action和异步Action
// 定义Thunk类型

export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
