import {
    Controller,
    type Control,
    type FieldValues,
    type Path,
    type RegisterOptions,
} from 'react-hook-form'
import type { InputProps } from 'antd'
import { Input,Typography } from 'antd'

type TextFieldProps<T extends FieldValues> = {
    inputProps: InputProps
    label: string
    name: Path<T>
    rules: RegisterOptions<T>
    control: Control<T>
    error: string
}

const TextField = <T extends FieldValues>(props: TextFieldProps<T>) => {
    
    //de-referencing props
    const { inputProps, name, rules, control, error } = props;
    inputProps.status = error ? "error" : undefined;

    const validationRules = {
        maxLength: {
            value: 250, message: "Text is too long"
        },
        pattern: {
            value: /^[\x20-\x9F\xA0-\xFF]*$/,
            message: "Invalid text format"
        },
        ...rules
    };
    
    return (
        <Controller
            control={control}
            name={name}
            rules={validationRules}
            render={({field: { onChange, value}}) => (
                <div className="form-item">
                    <Typography.Title level={5}>{props.label}</Typography.Title>
                    <Input
                        id={name}
                        size="small"
                        value={value}
                        onChange={(e) => onChange(e.target.value)}
                        placeholder='Filled'
                        {...inputProps}
                    />
                    <div className="error-message">{error}</div>
                </div>
            )}
        >

        </Controller>
    )
}

export default TextField;