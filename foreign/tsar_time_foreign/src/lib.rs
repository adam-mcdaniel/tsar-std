extern crate xmachine;
use xmachine::{Machine, Value};

extern crate time;
use time::now;

pub fn xasm_time_month(m: &mut Machine) {
    m.push(Value::number((now().tm_mon+1) as f64))
}

pub fn xasm_time_day(m: &mut Machine) {
    m.push(Value::number(now().tm_mday as f64))
}

pub fn xasm_time_year(m: &mut Machine) {
    m.push(Value::number((now().tm_year + 1900) as f64))
}

pub fn xasm_time_hour(m: &mut Machine) {
    m.push(Value::number(now().tm_hour as f64))
}

pub fn xasm_time_minute(m: &mut Machine) {
    m.push(Value::number(now().tm_min as f64))
}

pub fn xasm_time_second(m: &mut Machine) {
    m.push(Value::number(now().tm_sec as f64))
}

pub fn xasm_time_nanosecond(m: &mut Machine) {
    m.push(Value::number(now().tm_nsec as f64))
}
