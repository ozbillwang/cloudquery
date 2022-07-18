insert into aws_policy_results
select
  :execution_time,
  :'framework',
  :'check_id',
  'Ensure IAM password policy prevents password reuse',
  account_id,
  account_id,
  case when
    (password_reuse_prevention is null or password_reuse_prevention > 24)
    or policy_exists = FALSE
    then 'fail'
    else 'pass'
  end
from
    aws_iam_password_policies
