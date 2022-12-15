using System.Security.Claims;

namespace Janus.Models;

public class TokenRequestModel
{
    public string email { get; set; }
    public string[] roles { get; set; }
}