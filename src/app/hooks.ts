// 引入项目的state和dispatch类型
import type { RootState, AppDispatch } from './store';

// 定义Hook 在整个应用中使用
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
